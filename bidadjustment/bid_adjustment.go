package bidadjustment

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/holiman/uint256"
	"github.com/rs/zerolog"
)

func AdjustBlock(adjustmentData *AdjustmentData, gasFee uint64, gasUsed uint64, cumulativeGasUsed uint64, transferAmount uint64, adjustedPaymentAmount uint64, numTxs uint64, adjustedPaymentTx *types.Transaction, log *zerolog.Logger, isEOA bool, logs []*types.Log, rootNode *RootNode, builderState, feeRecipientState, payerState *types.StateAccount) (common.Hash, common.Hash, common.Hash, error) {
	stateRoot, err := adjustStateRoot(adjustmentData, gasFee, gasUsed, transferAmount, adjustedPaymentAmount, rootNode, builderState, feeRecipientState, payerState)
	if err != nil {
		log.Error().Err(err).Interface("adjustmentData", *adjustmentData).Msg("failed to adjust state root")
		return common.Hash{}, common.Hash{}, common.Hash{}, err
	}

	txRoot, err := adjustTxRoot(adjustmentData, numTxs, adjustedPaymentTx)
	if err != nil {
		log.Error().Err(err).Interface("adjustmentData", *adjustmentData).Uint64("numTxs", numTxs).Msg("failed to adjust tx root")
		return common.Hash{}, common.Hash{}, common.Hash{}, err
	}

	receiptRoot := adjustmentData.ReceiptsRoot
	if !isEOA {
		receiptRoot, _, _, err = adjustReceiptRoot(adjustmentData, adjustedPaymentTx, numTxs, stateRoot.Bytes(), cumulativeGasUsed, logs)
		if err != nil {
			return common.Hash{}, common.Hash{}, common.Hash{}, err
		}
	}
	return stateRoot, txRoot, receiptRoot, nil
}

func GetStateValueNodes(adjustmentData *AdjustmentData) (*RootNode, *types.StateAccount, *types.StateAccount, *types.StateAccount, error) {
	var (
		err       error
		valueByte []byte
		rootNode  RootNode
	)
	keys := [][]byte{
		crypto.Keccak256Hash(adjustmentData.BuilderAddress[:]).Bytes(),
		crypto.Keccak256Hash(adjustmentData.FeeRecipientAddress[:]).Bytes(),
		crypto.Keccak256Hash(adjustmentData.FeePayerAddress[:]).Bytes(),
	}
	proofSet := []*trienode.ProofSet{
		convertStateToTrienode(adjustmentData.BuilderProof).Set(),
		convertStateToTrienode(adjustmentData.FeeRecipientProof).Set(),
		convertStateToTrienode(adjustmentData.FeePayerProof).Set(),
	}

	valueNodes := [][]byte{}

	stateRoot := common.Hash(adjustmentData.StateRoot)
	for i := range keys {
		rootNode.Node, valueByte, err = proofToPath(stateRoot, rootNode.Node, keys[i], proofSet[i], false)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		valueNodes = append(valueNodes, valueByte)
	}
	fmt.Println("valueNodes[0]", hex.EncodeToString(valueNodes[0]))
	newBuilderState, err := valueNodeToAccount(valueNodes[0])
	if err != nil {
		return nil, nil, nil, nil, err
	}
	newFeeRecipientState, err := valueNodeToAccount(valueNodes[1])
	if err != nil {
		return nil, nil, nil, nil, err
	}

	newPayerState, err := valueNodeToAccount(valueNodes[2])
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &rootNode, &newBuilderState, &newFeeRecipientState, &newPayerState, nil
}

func adjustStateRoot(adjustmentData *AdjustmentData, gasFee uint64, gasUsed uint64, transferAmount uint64, adjustedPaymentAmount uint64, rootNode *RootNode, builderState, feeRecipientState, payerState *types.StateAccount) (common.Hash, error) {
	keys := [][]byte{
		crypto.Keccak256Hash(adjustmentData.BuilderAddress[:]).Bytes(),
		crypto.Keccak256Hash(adjustmentData.FeeRecipientAddress[:]).Bytes(),
		crypto.Keccak256Hash(adjustmentData.FeePayerAddress[:]).Bytes(),
	}

	adjustedValueNodes, err := adjustStateValueNodes(gasFee, gasUsed, transferAmount, adjustedPaymentAmount, builderState, feeRecipientState, payerState)
	if err != nil {
		return common.Hash{}, err
	}

	for i, key := range keys {
		k := keybytesToHex(key)
		valueNode := valueNode(adjustedValueNodes[i])
		_, updatedNode, err := rootNode.updateNode(rootNode.Node, nil, k, valueNode)
		if err != nil {
			return common.Hash{}, err
		}
		rootNode.Node = updatedNode
	}
	hash, _ := rootNode.Hash()

	return common.BytesToHash(hash.(hashNode)), nil
}

func adjustStateValueNodes(gasFee uint64, gasUsed uint64, transferAmount uint64, adjustedPaymentAmount uint64, oldBuilderState, oldFeeRecipientState, oldPayerState *types.StateAccount) ([][]byte, error) {
	gasCost := uint256.NewInt(gasUsed * gasFee)
	payment := uint256.NewInt(transferAmount)
	newBalance := uint256.NewInt(0)
	newBalance.Add(oldBuilderState.Balance, payment)
	newBalance.Add(newBalance, gasCost)
	newBuilderState := types.StateAccount{
		Nonce:    oldBuilderState.Nonce - 1,
		Balance:  newBalance,
		Root:     oldBuilderState.Root,
		CodeHash: oldBuilderState.CodeHash,
	}
	newBalance = uint256.NewInt(0)
	if oldFeeRecipientState.Balance.Cmp(payment) < 0 {
		return nil, fmt.Errorf("insufficient recipient balance: %d, payment: %d", oldFeeRecipientState.Balance, payment)
	}
	newBalance.Sub(oldFeeRecipientState.Balance, payment)
	adjustedPayment := uint256.NewInt(adjustedPaymentAmount)
	newBalance.Add(newBalance, adjustedPayment)
	newFeeRecipientState := types.StateAccount{
		Nonce:    oldFeeRecipientState.Nonce,
		Balance:  newBalance,
		Root:     oldFeeRecipientState.Root,
		CodeHash: oldFeeRecipientState.CodeHash,
	}

	newBalance = uint256.NewInt(0)
	cost := uint256.NewInt(0)
	cost.Add(gasCost, adjustedPayment)

	//Check collateral wallet has enough funds
	if oldPayerState.Balance.Cmp(cost) < 0 {
		return nil, fmt.Errorf("insufficient payer balance: %d, cost: %d", oldPayerState.Balance.Uint64(), cost.Uint64())
	}
	newBalance.Sub(oldPayerState.Balance, cost)
	newPayerState := types.StateAccount{
		Nonce:    oldPayerState.Nonce + 1,
		Balance:  newBalance,
		Root:     oldPayerState.Root,
		CodeHash: oldPayerState.CodeHash,
	}

	return [][]byte{
		encodeStateRLP(newBuilderState),
		encodeStateRLP(newFeeRecipientState),
		encodeStateRLP(newPayerState),
	}, nil
}

func convertStateToTrienode(stateProofs [][]byte) *trienode.ProofList {
	trienodeProof := make(trienode.ProofList, len(stateProofs))
	for i, proof := range stateProofs {
		trienodeProof[i] = proof // Adjust if the types differ
	}
	return &trienodeProof
}

func adjustTxRoot(adjustmentData *AdjustmentData, numTxs uint64, adjustedPaymentTx *types.Transaction) (common.Hash, error) {
	var (
		err      error
		rootNode RootNode
	)
	txList := types.Transactions{adjustedPaymentTx}
	encodedAdjustedPaymentTx := encodeListForDerive(txList, 0)

	transactionKey, err := rlp.EncodeToBytes(uint(numTxs - 1))
	if err != nil {
		return common.Hash{}, err
	}

	txRoot := common.Hash(adjustmentData.TransactionsRoot)
	proofDb := convertStateToTrienode(adjustmentData.PlaceholderTxProof).Set()
	rootNode.Node, _, err = proofToPath(txRoot, rootNode.Node, transactionKey, proofDb, false)
	if err != nil {
		return common.Hash{}, err
	}
	k := keybytesToHex(transactionKey)
	valueTrieNode := valueNode(encodedAdjustedPaymentTx)
	_, updatedNode, err := rootNode.updateNode(rootNode.Node, nil, k, valueTrieNode)
	if err != nil {
		return common.Hash{}, err
	}
	rootNode.Node = updatedNode

	hash, _ := rootNode.Hash()

	return common.BytesToHash(hash.(hashNode)), nil
}

func adjustReceiptRoot(adjustmentData *AdjustmentData, adjustedPaymentTx *types.Transaction, numTxs uint64, stateRoot []byte, cumulativeGasUsed uint64, logs []*types.Log) (common.Hash, valueNode, valueNode, error) {
	var (
		err          error
		oldValueNode []byte
		rootNode     RootNode
	)
	adjustedPaymentReceipt := &types.Receipt{
		Type:              adjustedPaymentTx.Type(),
		PostState:         []byte{},
		Status:            types.ReceiptStatusSuccessful,
		CumulativeGasUsed: cumulativeGasUsed, //Careful with gas refunds, shouldn't include gas refunds
		Logs:              logs,
	}
	adjustedPaymentReceipt.Bloom = types.CreateBloom(adjustedPaymentReceipt)

	receiptList := types.Receipts{adjustedPaymentReceipt}
	encodedAdjustedPaymentReceipt := encodeListForDerive(receiptList, 0)
	receiptKey, err := rlp.EncodeToBytes(uint(numTxs - 1))
	if err != nil {
		return common.Hash{}, nil, nil, err
	}

	receiptRoot := common.Hash(adjustmentData.ReceiptsRoot)
	proofDb := convertStateToTrienode(adjustmentData.PlaceholderReceiptProof).Set()

	rootNode.Node, oldValueNode, err = proofToPath(receiptRoot, rootNode.Node, receiptKey, proofDb, false)
	if err != nil {
		return common.Hash{}, nil, nil, err
	}

	k := keybytesToHex(receiptKey)
	newValueTrieNode := valueNode(encodedAdjustedPaymentReceipt)
	oldValueTrieNode := valueNode(oldValueNode)
	_, updatedNode, err := rootNode.updateNode(rootNode.Node, nil, k, newValueTrieNode)
	if err != nil {
		return common.Hash{}, nil, nil, err
	}
	rootNode.Node = updatedNode

	hash, _ := rootNode.Hash()

	return common.BytesToHash(hash.(hashNode)), oldValueTrieNode, newValueTrieNode, nil
}

func valueNodeToAccount(valueNode []byte) (types.StateAccount, error) {
	buf := bytes.NewBuffer(valueNode)
	stateFromProof := types.StateAccount{}
	err := rlp.Decode(buf, &stateFromProof)
	return stateFromProof, err
}

func encodeStateRLP(state types.StateAccount) []byte {
	var buf bytes.Buffer
	rlp.Encode(&buf, &state) //nolint:errcheck
	encoded := buf.Bytes()
	return encoded
}

func encodeListForDerive(list types.DerivableList, i int) []byte {
	// buf.Reset()
	w := new(bytes.Buffer)

	list.EncodeIndex(i, w)
	// It's really unfortunate that we need to do perform this copy.
	// StackTrie holds onto the values until Hash is called, so the values
	// written to it must not alias.
	return common.CopyBytes(w.Bytes())
}
