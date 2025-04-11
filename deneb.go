package relay_grpc

import (
	"fmt"

	apiDeneb "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	consensusspec "github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	consensus "github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
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

type SignedHeaderSubmissionDeneb struct {
	URL       string                  `json:"url"`
	Message   HeaderSubmissionDenebV2 `json:"message"`
	Signature phase0.BLSSignature     `json:"signature"`
}

type HeaderSubmissionDenebV2 struct {
	BidTrace               *v1.BidTrace                  `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader `json:"execution_payload_header"`
	Commitments            []deneb.KZGCommitment         `json:"commitments"`
}

func ProtoRequestToDenebHeaderSubmission(header *StreamHeaderResponse) (*SignedHeaderSubmissionDeneb, error) {
	bidTrace := header.BidTrace
	value, err := uint256.FromHex(bidTrace.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deneb block value %s to uint256: %s", bidTrace.Value, err.Error())
	}
	commitments := make([]deneb.KZGCommitment, len(header.Commitments))
	for i, commitment := range header.Commitments {
		copy(commitments[i][:], commitment)
	}
	signature := b96(header.Signature)

	return &SignedHeaderSubmissionDeneb{
		URL: "",
		Message: HeaderSubmissionDenebV2{
			BidTrace: &v1.BidTrace{
				Slot:                 bidTrace.Slot,
				ParentHash:           b32(bidTrace.ParentHash),
				BlockHash:            b32(bidTrace.BlockHash),
				BuilderPubkey:        b48(bidTrace.BuilderPubkey),
				ProposerPubkey:       b48(bidTrace.ProposerPubkey),
				ProposerFeeRecipient: b20(bidTrace.ProposerFeeRecipient),
				GasLimit:             bidTrace.GasLimit,
				GasUsed:              bidTrace.GasUsed,
				Value:                value,
			},
			ExecutionPayloadHeader: &consensus.ExecutionPayloadHeader{
				ParentHash:       b32(header.ExecutionPayloadHeader.ParentHash),
				StateRoot:        b32(header.ExecutionPayloadHeader.StateRoot),
				ReceiptsRoot:     b32(header.ExecutionPayloadHeader.ReceiptsRoot),
				LogsBloom:        b256(header.ExecutionPayloadHeader.LogsBloom),
				PrevRandao:       b32(header.ExecutionPayloadHeader.PrevRandao),
				BaseFeePerGas:    byteSliceToUint256Int(header.ExecutionPayloadHeader.BaseFeePerGas),
				FeeRecipient:     b20(header.ExecutionPayloadHeader.FeeRecipient),
				BlockHash:        b32(header.ExecutionPayloadHeader.BlockHash),
				ExtraData:        header.ExecutionPayloadHeader.ExtraData,
				BlockNumber:      header.ExecutionPayloadHeader.BlockNumber,
				GasLimit:         header.ExecutionPayloadHeader.GasLimit,
				Timestamp:        header.ExecutionPayloadHeader.Timestamp,
				GasUsed:          header.ExecutionPayloadHeader.GasUsed,
				TransactionsRoot: b32(header.ExecutionPayloadHeader.TransactionsRoot),
				WithdrawalsRoot:  b32(header.ExecutionPayloadHeader.WithdrawalsRoot),
				BlobGasUsed:      header.ExecutionPayloadHeader.BlobGasUsed,
				ExcessBlobGas:    header.ExecutionPayloadHeader.ExcessBlobGas,
			},
		},
		Signature: signature,
	}, nil
}

func DenebBlockRequestToHeaderSubmissionProtoRequest(block *apiDeneb.SubmitBlockRequest, transactionsRoot []byte, withdrawalsRoot []byte) (*BidTrace, *ExecutionPayloadHeader, [][]byte, []byte) {
	commitments := make([][]byte, len(block.BlobsBundle.Commitments))

	for i, commitment := range block.BlobsBundle.Commitments {
		commitments[i] = commitment[:]
	}

	return &BidTrace{
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
		}, &ExecutionPayloadHeader{
			ParentHash:       block.ExecutionPayload.ParentHash[:],
			StateRoot:        block.ExecutionPayload.StateRoot[:],
			ReceiptsRoot:     block.ExecutionPayload.ReceiptsRoot[:],
			LogsBloom:        block.ExecutionPayload.LogsBloom[:],
			PrevRandao:       block.ExecutionPayload.PrevRandao[:],
			BaseFeePerGas:    uint256ToIntToByteSlice(block.ExecutionPayload.BaseFeePerGas),
			FeeRecipient:     block.ExecutionPayload.FeeRecipient[:],
			BlockHash:        block.ExecutionPayload.BlockHash[:],
			ExtraData:        block.ExecutionPayload.ExtraData,
			BlockNumber:      block.ExecutionPayload.BlockNumber,
			GasLimit:         block.ExecutionPayload.GasLimit,
			Timestamp:        block.ExecutionPayload.Timestamp,
			GasUsed:          block.ExecutionPayload.GasUsed,
			TransactionsRoot: transactionsRoot,
			WithdrawalsRoot:  withdrawalsRoot,
			BlobGasUsed:      block.ExecutionPayload.BlobGasUsed,
			ExcessBlobGas:    block.ExecutionPayload.ExcessBlobGas,
		},
		commitments,
		block.Signature[:]
}
