package relay_grpc

import (
	"encoding/json"

	apiv1deneb "github.com/attestantio/go-eth2-client/api/v1/deneb"

	apiDeneb "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	consensus "github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

func ExecutionPayloadToProtoExecutionPayloadAndBlobsBundle(executionPayload *apiDeneb.ExecutionPayloadAndBlobsBundle) *ExecutionPayloadAndBlobsBundle {
	transactions := make([]*Transaction, len(executionPayload.ExecutionPayload.Transactions))
	for i, tx := range executionPayload.ExecutionPayload.Transactions {
		transactions[i] = &Transaction{
			RawData: tx,
		}
	}
	withdrawals := make([]*Withdrawal, len(executionPayload.ExecutionPayload.Withdrawals))
	for i, withdrawal := range executionPayload.ExecutionPayload.Withdrawals {
		withdrawals[i] = &Withdrawal{
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Index:          uint64(withdrawal.Index),
			Amount:         uint64(withdrawal.Amount),
			Address:        withdrawal.Address[:],
		}
	}
	ExecutionPayloadUncompressed := &ExecutionPayloadUncompressed{
		ParentHash:    executionPayload.ExecutionPayload.ParentHash[:],
		StateRoot:     executionPayload.ExecutionPayload.StateRoot[:],
		ReceiptsRoot:  executionPayload.ExecutionPayload.ReceiptsRoot[:],
		LogsBloom:     executionPayload.ExecutionPayload.LogsBloom[:],
		PrevRandao:    executionPayload.ExecutionPayload.PrevRandao[:],
		BaseFeePerGas: uint256ToIntToByteSlice(executionPayload.ExecutionPayload.BaseFeePerGas),
		FeeRecipient:  executionPayload.ExecutionPayload.FeeRecipient[:],
		BlockHash:     executionPayload.ExecutionPayload.BlockHash[:],
		ExtraData:     executionPayload.ExecutionPayload.ExtraData,
		BlockNumber:   executionPayload.ExecutionPayload.BlockNumber,
		GasLimit:      executionPayload.ExecutionPayload.GasLimit,
		Timestamp:     executionPayload.ExecutionPayload.Timestamp,
		GasUsed:       executionPayload.ExecutionPayload.GasUsed,
		Transactions:  transactions,
		Withdrawals:   withdrawals,
		BlobGasUsed:   executionPayload.ExecutionPayload.BlobGasUsed,
		ExcessBlobGas: executionPayload.ExecutionPayload.ExcessBlobGas,
	}
	BlobsBundle := convertBlobBundleToProto(executionPayload.BlobsBundle)
	return &ExecutionPayloadAndBlobsBundle{
		ExecutionPayload: ExecutionPayloadUncompressed,
		BlobsBundle:      BlobsBundle,
	}
}
func ExecutionPayloadToProtoEnrichBlockRequest(uuid string,
	executionPayload *apiDeneb.ExecutionPayloadAndBlobsBundle,
	bidTrace v1.BidTrace, parentBeaconRoot phase0.Root) EnrichBlockRequest {
	BidTrace := &BidTrace{
		Slot:                 bidTrace.Slot,
		ParentHash:           bidTrace.ParentHash[:],
		BlockHash:            bidTrace.BlockHash[:],
		BuilderPubkey:        bidTrace.BuilderPubkey[:],
		ProposerPubkey:       bidTrace.ProposerPubkey[:],
		ProposerFeeRecipient: bidTrace.ProposerFeeRecipient[:],
		GasLimit:             bidTrace.GasLimit,
		GasUsed:              bidTrace.GasUsed,
		Value:                bidTrace.Value.Hex(),
		BlobGasUsed:          executionPayload.ExecutionPayload.BlobGasUsed,
		ExcessBlobGas:        executionPayload.ExecutionPayload.ExcessBlobGas,
	}
	execPayloadAndBlobsBundle := ExecutionPayloadToProtoExecutionPayloadAndBlobsBundle(executionPayload)
	return EnrichBlockRequest{
		Uuid:                           uuid,
		ExecutionPayloadAndBlobsBundle: execPayloadAndBlobsBundle,
		BidTrace:                       BidTrace,
		ParentBeaconRoot:               parentBeaconRoot[:],
	}
}

func ProtoEnrichBlockResponseToExecutionPayloadHeader(enrichBlockResponse *EnrichBlockResponse) (string,
	consensus.ExecutionPayloadHeader, //enriched execution payload
	[]consensus.KZGCommitment, //enriched commitments
	uint256.Int, //enriched value
) {
	uuid := enrichBlockResponse.GetUuid()
	commitments := make([]consensus.KZGCommitment, len(enrichBlockResponse.GetKzgCommitment()))
	for i, commitment := range enrichBlockResponse.KzgCommitment {
		copy(commitments[i][:], commitment)
	}
	protoExecutionPayloadHeader := enrichBlockResponse.GetExecutionPayloadHeader()
	newValue := enrichBlockResponse.GetValue()

	executionPayloadHeader := &consensus.ExecutionPayloadHeader{
		ParentHash:       b32(protoExecutionPayloadHeader.ParentHash),
		StateRoot:        b32(protoExecutionPayloadHeader.StateRoot),
		ReceiptsRoot:     b32(protoExecutionPayloadHeader.ReceiptsRoot),
		LogsBloom:        b256(protoExecutionPayloadHeader.LogsBloom),
		PrevRandao:       b32(protoExecutionPayloadHeader.PrevRandao),
		BaseFeePerGas:    byteSliceToUint256Int(protoExecutionPayloadHeader.BaseFeePerGas),
		FeeRecipient:     b20(protoExecutionPayloadHeader.FeeRecipient),
		BlockHash:        b32(protoExecutionPayloadHeader.BlockHash),
		ExtraData:        protoExecutionPayloadHeader.ExtraData,
		BlockNumber:      protoExecutionPayloadHeader.BlockNumber,
		GasLimit:         protoExecutionPayloadHeader.GasLimit,
		Timestamp:        protoExecutionPayloadHeader.Timestamp,
		GasUsed:          protoExecutionPayloadHeader.GasUsed,
		TransactionsRoot: b32(protoExecutionPayloadHeader.TransactionsRoot),
		WithdrawalsRoot:  b32(protoExecutionPayloadHeader.WithdrawalsRoot),
		BlobGasUsed:      protoExecutionPayloadHeader.BlobGasUsed,
		ExcessBlobGas:    protoExecutionPayloadHeader.ExcessBlobGas,
	}
	return uuid, *executionPayloadHeader, commitments, uint256.Int{newValue}
}

func SignedBeaconBlockToProtoGetEnrichedPayloadRequest(signedBeaconBlock *apiv1deneb.SignedBlindedBeaconBlock) *GetEnrichedPayloadRequest {
	signedBeaconBlockBytes, _ := json.Marshal(signedBeaconBlock.Message)
	return &GetEnrichedPayloadRequest{
		Message:   signedBeaconBlockBytes,
		Signature: signedBeaconBlock.Signature[:],
	}
}

func ProtoGetEnrichedPayloadResponseToExecutionPayload(executionPayloadAndBlobsBundle *ExecutionPayloadAndBlobsBundle) *apiDeneb.ExecutionPayloadAndBlobsBundle {
	transactions := make([]bellatrix.Transaction, len(executionPayloadAndBlobsBundle.ExecutionPayload.Transactions))
	for index, tx := range executionPayloadAndBlobsBundle.ExecutionPayload.Transactions {
		transactions[index] = tx.RawData
	}

	// Withdrawal is defined in capella spec
	// https://github.com/attestantio/go-eth2-client/blob/21f7dd480fed933d8e0b1c88cee67da721c80eb2/spec/deneb/executionpayload.go#L42
	withdrawals := make([]*capella.Withdrawal, len(executionPayloadAndBlobsBundle.ExecutionPayload.Withdrawals))
	for index, withdrawal := range executionPayloadAndBlobsBundle.ExecutionPayload.Withdrawals {
		withdrawals[index] = &capella.Withdrawal{
			ValidatorIndex: phase0.ValidatorIndex(withdrawal.ValidatorIndex),
			Index:          capella.WithdrawalIndex(withdrawal.Index),
			Amount:         phase0.Gwei(withdrawal.Amount),
			Address:        b20(withdrawal.Address),
		}
	}
	blobsBundle := &apiDeneb.BlobsBundle{
		Commitments: make([]consensus.KZGCommitment, len(executionPayloadAndBlobsBundle.BlobsBundle.Commitments)),
		Proofs:      make([]consensus.KZGProof, len(executionPayloadAndBlobsBundle.BlobsBundle.Proofs)),
		Blobs:       make([]consensus.Blob, len(executionPayloadAndBlobsBundle.BlobsBundle.Blobs)),
	}
	for index, commitment := range executionPayloadAndBlobsBundle.BlobsBundle.Commitments {
		copy(blobsBundle.Commitments[index][:], commitment)
	}

	for index, proof := range executionPayloadAndBlobsBundle.BlobsBundle.Proofs {
		copy(blobsBundle.Proofs[index][:], proof)
	}

	for index, blob := range executionPayloadAndBlobsBundle.BlobsBundle.Blobs {
		copy(blobsBundle.Blobs[index][:], blob)
	}
	result := &apiDeneb.ExecutionPayloadAndBlobsBundle{
		ExecutionPayload: &consensus.ExecutionPayload{
			ParentHash:    b32(executionPayloadAndBlobsBundle.ExecutionPayload.ParentHash),
			StateRoot:     b32(executionPayloadAndBlobsBundle.ExecutionPayload.StateRoot),
			ReceiptsRoot:  b32(executionPayloadAndBlobsBundle.ExecutionPayload.ReceiptsRoot),
			LogsBloom:     b256(executionPayloadAndBlobsBundle.ExecutionPayload.LogsBloom),
			PrevRandao:    b32(executionPayloadAndBlobsBundle.ExecutionPayload.PrevRandao),
			BaseFeePerGas: byteSliceToUint256Int(executionPayloadAndBlobsBundle.ExecutionPayload.BaseFeePerGas),
			FeeRecipient:  b20(executionPayloadAndBlobsBundle.ExecutionPayload.FeeRecipient),
			BlockHash:     b32(executionPayloadAndBlobsBundle.ExecutionPayload.BlockHash),
			ExtraData:     executionPayloadAndBlobsBundle.ExecutionPayload.ExtraData,
			BlockNumber:   executionPayloadAndBlobsBundle.ExecutionPayload.BlockNumber,
			GasLimit:      executionPayloadAndBlobsBundle.ExecutionPayload.GasLimit,
			Timestamp:     executionPayloadAndBlobsBundle.ExecutionPayload.Timestamp,
			GasUsed:       executionPayloadAndBlobsBundle.ExecutionPayload.GasUsed,
			Transactions:  transactions,
			Withdrawals:   withdrawals,
			BlobGasUsed:   executionPayloadAndBlobsBundle.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas: executionPayloadAndBlobsBundle.ExecutionPayload.ExcessBlobGas,
		},
		BlobsBundle: blobsBundle,
	}
	return result
}

func ExecutionPayloadHeaderToProtoResponse(uuid string, message *apiDeneb.BuilderBid) *EnrichBlockResponse {
	executionPayload := message.Header
	executionPayloadHeader := &ExecutionPayloadHeader{
		ParentHash:       executionPayload.ParentHash[:],
		StateRoot:        executionPayload.StateRoot[:],
		ReceiptsRoot:     executionPayload.ReceiptsRoot[:],
		LogsBloom:        executionPayload.LogsBloom[:],
		PrevRandao:       executionPayload.PrevRandao[:],
		BaseFeePerGas:    uint256ToIntToByteSlice(executionPayload.BaseFeePerGas),
		FeeRecipient:     executionPayload.FeeRecipient[:],
		BlockHash:        executionPayload.BlockHash[:],
		ExtraData:        executionPayload.ExtraData,
		BlockNumber:      executionPayload.BlockNumber,
		GasLimit:         executionPayload.GasLimit,
		Timestamp:        executionPayload.Timestamp,
		GasUsed:          executionPayload.GasUsed,
		TransactionsRoot: executionPayload.TransactionsRoot[:],
		WithdrawalsRoot:  executionPayload.WithdrawalsRoot[:],
		BlobGasUsed:      executionPayload.BlobGasUsed,
		ExcessBlobGas:    executionPayload.ExcessBlobGas,
	}
	commitments := make([][]byte, 0)
	for _, commitment := range message.BlobKZGCommitments {
		commitments = append(commitments, commitment[:])
	}

	enrichedBlockResponse := &EnrichBlockResponse{
		Uuid:                   uuid,
		ExecutionPayloadHeader: executionPayloadHeader,
		KzgCommitment:          commitments,
		Value:                  message.Value.Uint64(),
	}
	return enrichedBlockResponse
}
