package relay_grpc

import (
	"fmt"

	apiDeneb "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	consensusspec "github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	consensus "github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

func DenebRequestToProtoRequest(block *apiDeneb.SubmitBlockRequest) *SubmitBlockRequest {
	transactions := []*CompressTx{}
	for _, tx := range block.ExecutionPayload.Transactions {
		transactions = append(transactions, &CompressTx{
			RawData: tx,
			ShortID: 0,
		})
	}

	withdrawals := []*Withdrawal{}
	for _, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals = append(withdrawals, &Withdrawal{
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Index:          uint64(withdrawal.Index),
			Amount:         uint64(withdrawal.Amount),
			Address:        withdrawal.Address[:],
		})
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
	withdrawals := []*Withdrawal{}
	for _, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals = append(withdrawals, &Withdrawal{
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Index:          uint64(withdrawal.Index),
			Amount:         uint64(withdrawal.Amount),
			Address:        withdrawal.Address[:],
		})
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
	transactions := []bellatrix.Transaction{}
	for _, tx := range block.ExecutionPayload.Transactions {
		transactions = append(transactions, tx.RawData)
	}

	// Withdrawal is defined in capella spec
	// https://github.com/attestantio/go-eth2-client/blob/21f7dd480fed933d8e0b1c88cee67da721c80eb2/spec/deneb/executionpayload.go#L42
	withdrawals := []*capella.Withdrawal{}
	for _, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals = append(withdrawals, &capella.Withdrawal{
			ValidatorIndex: phase0.ValidatorIndex(withdrawal.ValidatorIndex),
			Index:          capella.WithdrawalIndex(withdrawal.Index),
			Amount:         phase0.Gwei(withdrawal.Amount),
			Address:        b20(withdrawal.Address),
		})
	}

	// BlobsBundle
	blobsBundle := &apiDeneb.BlobsBundle{
		Commitments: []consensus.KZGCommitment{},
		Proofs:      []consensus.KZGProof{},
		Blobs:       []consensus.Blob{},
	}
	for _, commitment := range block.BlobsBundle.Commitments {
		blobsBundle.Commitments = append(blobsBundle.Commitments, consensus.KZGCommitment(commitment))
	}

	for _, proof := range block.BlobsBundle.Proofs {
		blobsBundle.Proofs = append(blobsBundle.Proofs, consensus.KZGProof(proof))
	}

	for _, blob := range block.BlobsBundle.Blobs {
		blobsBundle.Blobs = append(blobsBundle.Blobs, consensus.Blob(blob))
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
		Commitments: [][]byte{},
		Proofs:      [][]byte{},
		Blobs:       [][]byte{},
	}

	for i := range blobBundle.Commitments {
		protoBlobsBundle.Commitments = append(protoBlobsBundle.Commitments, blobBundle.Commitments[i][:])
	}

	for i := range blobBundle.Proofs {
		protoBlobsBundle.Proofs = append(protoBlobsBundle.Proofs, blobBundle.Proofs[i][:])
	}

	for i := range blobBundle.Blobs {
		protoBlobsBundle.Blobs = append(protoBlobsBundle.Blobs, blobBundle.Blobs[i][:])
	}

	return protoBlobsBundle
}
