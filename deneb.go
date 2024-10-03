package relay_grpc

import (
	"encoding/json"
	"fmt"

	builderApi "github.com/attestantio/go-builder-client/api"
	apiDeneb "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	apiv1deneb "github.com/attestantio/go-eth2-client/api/v1/deneb"
	consensusspec "github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	consensus "github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/flashbots/go-boost-utils/utils"
	"github.com/holiman/uint256"
)

func DenebRequestToProtoRequest(block *apiDeneb.SubmitBlockRequest) *SubmitBlockRequest {
	transactions := make([]*CompressTx, len(block.ExecutionPayload.Transactions))
	for i, tx := range block.ExecutionPayload.Transactions {
		transactions[i] = &CompressTx{
			RawData: tx,
			ShortID: 0,
		}
	}

	withdrawals := make([]*Withdrawal, len(block.ExecutionPayload.Withdrawals))
	for i, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals[i] = &Withdrawal{
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Index:          uint64(withdrawal.Index),
			Amount:         uint64(withdrawal.Amount),
			Address:        withdrawal.Address[:],
		}
	}

	return &SubmitBlockRequest{
		Version: uint64(consensusspec.DataVersionDeneb),
		BidTrace: &BidTrace{
			Slot:                 block.Message.Slot,
			ParentHash:           block.Message.ParentHash[:],
			BlockHash:            block.Message.BlockHash[:],
			BuilderPubkey:        block.Message.BuilderPubkey[:],
			ProposerPubkey:       block.Message.ProposerPubkey[:],
			ProposerFeeRecipient: block.Message.ProposerFeeRecipient[:],
			GasLimit:             block.Message.GasLimit,
			GasUsed:              block.Message.GasUsed,
			Value:                block.Message.Value.Hex(),
			BlobGasUsed:          block.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas:        block.ExecutionPayload.ExcessBlobGas,
		},
		ExecutionPayload: &ExecutionPayload{
			ParentHash:    block.ExecutionPayload.ParentHash[:],
			StateRoot:     block.ExecutionPayload.StateRoot[:],
			ReceiptsRoot:  block.ExecutionPayload.ReceiptsRoot[:],
			LogsBloom:     block.ExecutionPayload.LogsBloom[:],
			PrevRandao:    block.ExecutionPayload.PrevRandao[:],
			BaseFeePerGas: uint256ToIntToByteSlice(block.ExecutionPayload.BaseFeePerGas),
			FeeRecipient:  block.ExecutionPayload.FeeRecipient[:],
			BlockHash:     block.ExecutionPayload.BlockHash[:],
			ExtraData:     block.ExecutionPayload.ExtraData,
			BlockNumber:   block.ExecutionPayload.BlockNumber,
			GasLimit:      block.ExecutionPayload.GasLimit,
			Timestamp:     block.ExecutionPayload.Timestamp,
			GasUsed:       block.ExecutionPayload.GasUsed,
			Transactions:  transactions,
			Withdrawals:   withdrawals,
			BlobGasUsed:   block.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas: block.ExecutionPayload.ExcessBlobGas,
		},
		BlobsBundle: convertBlobBundleToProto(block.BlobsBundle),
		Signature:   block.Signature[:],
	}
}

// DenebRequestToProtoRequest converts a Deneb request to a SubmitBlockRequest.
func DenebRequestToProtoRequestWithShortIDs(block *apiDeneb.SubmitBlockRequest, compressTxs []*CompressTx) *SubmitBlockRequest {
	withdrawals := make([]*Withdrawal, len(block.ExecutionPayload.Withdrawals))
	for i, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals[i] = &Withdrawal{
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Index:          uint64(withdrawal.Index),
			Amount:         uint64(withdrawal.Amount),
			Address:        withdrawal.Address[:],
		}
	}

	return &SubmitBlockRequest{
		Version: uint64(consensusspec.DataVersionDeneb),
		BidTrace: &BidTrace{
			Slot:                 block.Message.Slot,
			ParentHash:           block.Message.ParentHash[:],
			BlockHash:            block.Message.BlockHash[:],
			BuilderPubkey:        block.Message.BuilderPubkey[:],
			ProposerPubkey:       block.Message.ProposerPubkey[:],
			ProposerFeeRecipient: block.Message.ProposerFeeRecipient[:],
			GasLimit:             block.Message.GasLimit,
			GasUsed:              block.Message.GasUsed,
			Value:                block.Message.Value.Hex(),
			BlobGasUsed:          block.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas:        block.ExecutionPayload.ExcessBlobGas,
		},
		ExecutionPayload: &ExecutionPayload{
			ParentHash:    block.ExecutionPayload.ParentHash[:],
			StateRoot:     block.ExecutionPayload.StateRoot[:],
			ReceiptsRoot:  block.ExecutionPayload.ReceiptsRoot[:],
			LogsBloom:     block.ExecutionPayload.LogsBloom[:],
			PrevRandao:    block.ExecutionPayload.PrevRandao[:],
			BaseFeePerGas: uint256ToIntToByteSlice(block.ExecutionPayload.BaseFeePerGas),
			FeeRecipient:  block.ExecutionPayload.FeeRecipient[:],
			BlockHash:     block.ExecutionPayload.BlockHash[:],
			ExtraData:     block.ExecutionPayload.ExtraData,
			BlockNumber:   block.ExecutionPayload.BlockNumber,
			GasLimit:      block.ExecutionPayload.GasLimit,
			Timestamp:     block.ExecutionPayload.Timestamp,
			GasUsed:       block.ExecutionPayload.GasUsed,
			Transactions:  compressTxs,
			Withdrawals:   withdrawals,
			BlobGasUsed:   block.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas: block.ExecutionPayload.ExcessBlobGas,
		},
		BlobsBundle: convertBlobBundleToProto(block.BlobsBundle),
		Signature:   block.Signature[:],
	}
}

func ProtoRequestToDenebRequest(block *SubmitBlockRequest) (*apiDeneb.SubmitBlockRequest, error) {
	transactions := make([]bellatrix.Transaction, len(block.ExecutionPayload.Transactions))
	for index, tx := range block.ExecutionPayload.Transactions {
		transactions[index] = tx.RawData
	}

	// Withdrawal is defined in capella spec
	// https://github.com/attestantio/go-eth2-client/blob/21f7dd480fed933d8e0b1c88cee67da721c80eb2/spec/deneb/executionpayload.go#L42
	withdrawals := make([]*capella.Withdrawal, len(block.ExecutionPayload.Withdrawals))
	for index, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals[index] = &capella.Withdrawal{
			ValidatorIndex: phase0.ValidatorIndex(withdrawal.ValidatorIndex),
			Index:          capella.WithdrawalIndex(withdrawal.Index),
			Amount:         phase0.Gwei(withdrawal.Amount),
			Address:        b20(withdrawal.Address),
		}
	}

	// BlobsBundle
	blobsBundle := &apiDeneb.BlobsBundle{
		Commitments: make([]consensus.KZGCommitment, len(block.BlobsBundle.Commitments)),
		Proofs:      make([]consensus.KZGProof, len(block.BlobsBundle.Proofs)),
		Blobs:       make([]consensus.Blob, len(block.BlobsBundle.Blobs)),
	}
	for index, commitment := range block.BlobsBundle.Commitments {
		copy(blobsBundle.Commitments[index][:], commitment)
	}

	for index, proof := range block.BlobsBundle.Proofs {
		copy(blobsBundle.Proofs[index][:], proof)
	}

	for index, blob := range block.BlobsBundle.Blobs {
		copy(blobsBundle.Blobs[index][:], blob)
	}

	value, err := uint256.FromHex(block.BidTrace.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deneb block value %s to uint256: %s", block.BidTrace.Value, err.Error())
	}

	return &apiDeneb.SubmitBlockRequest{
		Message: &v1.BidTrace{
			Slot:                 block.BidTrace.Slot,
			ParentHash:           b32(block.BidTrace.ParentHash),
			BlockHash:            b32(block.BidTrace.BlockHash),
			BuilderPubkey:        b48(block.BidTrace.BuilderPubkey),
			ProposerPubkey:       b48(block.BidTrace.ProposerPubkey),
			ProposerFeeRecipient: b20(block.BidTrace.ProposerFeeRecipient),
			GasLimit:             block.BidTrace.GasLimit,
			GasUsed:              block.BidTrace.GasUsed,
			Value:                value,
		},
		ExecutionPayload: &consensus.ExecutionPayload{
			ParentHash:    b32(block.ExecutionPayload.ParentHash),
			StateRoot:     b32(block.ExecutionPayload.StateRoot),
			ReceiptsRoot:  b32(block.ExecutionPayload.ReceiptsRoot),
			LogsBloom:     b256(block.ExecutionPayload.LogsBloom),
			PrevRandao:    b32(block.ExecutionPayload.PrevRandao),
			BaseFeePerGas: byteSliceToUint256Int(block.ExecutionPayload.BaseFeePerGas),
			FeeRecipient:  b20(block.ExecutionPayload.FeeRecipient),
			BlockHash:     b32(block.ExecutionPayload.BlockHash),
			ExtraData:     block.ExecutionPayload.ExtraData,
			BlockNumber:   block.ExecutionPayload.BlockNumber,
			GasLimit:      block.ExecutionPayload.GasLimit,
			Timestamp:     block.ExecutionPayload.Timestamp,
			GasUsed:       block.ExecutionPayload.GasUsed,
			Transactions:  transactions,
			Withdrawals:   withdrawals,
			BlobGasUsed:   block.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas: block.ExecutionPayload.ExcessBlobGas,
		},
		BlobsBundle: blobsBundle,
		Signature:   b96(block.Signature),
	}, nil
}

// Add Commitments, Proofs, Data to BlobsBundle
func convertBlobBundleToProto(blobBundle *apiDeneb.BlobsBundle) *BlobsBundle {
	protoBlobsBundle := &BlobsBundle{
		Commitments: make([][]byte, len(blobBundle.Commitments)),
		Proofs:      make([][]byte, len(blobBundle.Proofs)),
		Blobs:       make([][]byte, len(blobBundle.Blobs)),
	}

	for i := range blobBundle.Commitments {
		protoBlobsBundle.Commitments[i] = blobBundle.Commitments[i][:]
	}

	for i := range blobBundle.Proofs {
		protoBlobsBundle.Proofs[i] = blobBundle.Proofs[i][:]
	}

	for i := range blobBundle.Blobs {
		protoBlobsBundle.Blobs[i] = blobBundle.Blobs[i][:]
	}

	return protoBlobsBundle
}

func EnrichBlock(
	uuid string,
	executionPayload *apiDeneb.ExecutionPayloadAndBlobsBundle,
	bidTrace v1.BidTrace,
	parentBeaconRoot phase0.Root,
) (
	string,
	consensus.ExecutionPayloadHeader, //enriched execution payload
	[]consensus.KZGCommitment, //enriched commitments
	uint256.Int, //enriched value
) {
	versionedPayload := &builderApi.VersionedExecutionPayload{Version: consensusspec.DataVersionDeneb}
	versionedPayload.Deneb = executionPayload.ExecutionPayload
	payloadHeader, err := utils.PayloadToPayloadHeader(versionedPayload)
	if err != nil {
		return uuid, consensus.ExecutionPayloadHeader{}, nil, uint256.Int{}
	}
	var newValue uint256.Int
	newValue.Add(bidTrace.Value, uint256.NewInt(1))
	return uuid, *payloadHeader.Deneb, executionPayload.BlobsBundle.Commitments, newValue
}

func GetEnrichedPayload(
	apiv1deneb.SignedBlindedBeaconBlock, // commitment to enriched block
) *apiDeneb.ExecutionPayloadAndBlobsBundle {
	return nil
}

func GetEnrichedPayloadMock(
	signedBlindedBeaconBlock apiv1deneb.SignedBlindedBeaconBlock, // commitment to enriched block
	executionPayload *apiDeneb.ExecutionPayloadAndBlobsBundle,
) *apiDeneb.ExecutionPayloadAndBlobsBundle {

	return executionPayload
}

func ProtoRequestToDenebBidtracePayload(block *SubmitBlockRequest) (*BidtracePayload, error) {
	value, err := uint256.FromHex(block.BidTrace.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deneb block value %s to uint256: %s", block.BidTrace.Value, err.Error())
	}

	return &BidtracePayload{
		Message: &v1.BidTrace{
			Slot:                 block.BidTrace.Slot,
			ParentHash:           b32(block.BidTrace.ParentHash),
			BlockHash:            b32(block.BidTrace.BlockHash),
			BuilderPubkey:        b48(block.BidTrace.BuilderPubkey),
			ProposerPubkey:       b48(block.BidTrace.ProposerPubkey),
			ProposerFeeRecipient: b20(block.BidTrace.ProposerFeeRecipient),
			GasLimit:             block.BidTrace.GasLimit,
			GasUsed:              block.BidTrace.GasUsed,
			Value:                value,
		},
		ExecutionPayload: &BidTraceExecutionPayload{
			Timestamp: block.ExecutionPayload.Timestamp,
		},
		Signature: b96(block.Signature),
	}, nil
}

func ExecutionPayloadToProtoEnrichBlockRequest(uuid string,
	executionPayload *apiDeneb.ExecutionPayloadAndBlobsBundle,
	bidTrace v1.BidTrace, parentBeaconRoot phase0.Root) EnrichBlockRequest {
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
	return EnrichBlockRequest{
		Uuid: uuid,
		ExecutionPayloadAndBlobsBundle: &ExecutionPayloadAndBlobsBundle{
			ExecutionPayload: ExecutionPayloadUncompressed,
			BlobsBundle:      BlobsBundle,
		},
		BidTrace:         BidTrace,
		ParentBeaconRoot: parentBeaconRoot[:],
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
