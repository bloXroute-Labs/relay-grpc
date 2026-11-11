package relay_grpc

import (
	"fmt"

	apiFulu "github.com/attestantio/go-builder-client/api/fulu"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	consensusspec "github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/gloas"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

// GloasRequest is a builder's block submission for a Gloas slot, in go-eth2-client's
// types: the bid trace, the root of the beacon block the proposer's block builds on
// (EIP-4788's parent_beacon_block_root), which the builder signs with it, the Gloas
// payload, Fulu's blobs bundle (128 cell proofs per blob) and the Gloas execution requests.
// It holds what Helix's and Ultrasound's SignedBidSubmissionGloas holds, but flat: there
// the message nests the trace and the parent block root, which the builder signs
// together; here they are side by side, as the proto carries them. This type is never
// encoded or signed, only converted. go-builder-client has no Gloas submission, so the
// converters below use this, and the relay builds its own submission type from it.
type GloasRequest struct {
	Message           *v1.BidTrace
	ParentBlockRoot   phase0.Root
	ExecutionPayload  *gloas.ExecutionPayload
	BlobsBundle       *apiFulu.BlobsBundle
	ExecutionRequests *gloas.ExecutionRequests
	Signature         phase0.BLSSignature
}

// GloasRequestToProtoRequest converts a Gloas submission to a SubmitBlockRequest.
func GloasRequestToProtoRequest(block *GloasRequest) *SubmitBlockRequest {
	payload := block.ExecutionPayload
	transactions := make([]*CompressTx, len(payload.Transactions))
	for i, tx := range payload.Transactions {
		transactions[i] = &CompressTx{RawData: tx}
	}

	withdrawals := make([]*Withdrawal, len(payload.Withdrawals))
	for i, withdrawal := range payload.Withdrawals {
		withdrawals[i] = &Withdrawal{
			ValidatorIndex: uint64(withdrawal.ValidatorIndex),
			Index:          uint64(withdrawal.Index),
			Amount:         uint64(withdrawal.Amount),
			Address:        withdrawal.Address[:],
		}
	}

	return &SubmitBlockRequest{
		Version: uint64(consensusspec.DataVersionGloas),
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
			BlobGasUsed:          payload.BlobGasUsed,
			ExcessBlobGas:        payload.ExcessBlobGas,
		},
		ParentBlockRoot: block.ParentBlockRoot[:],
		ExecutionPayload: &ExecutionPayload{
			ParentHash:      payload.ParentHash[:],
			StateRoot:       payload.StateRoot[:],
			ReceiptsRoot:    payload.ReceiptsRoot[:],
			LogsBloom:       payload.LogsBloom[:],
			PrevRandao:      payload.PrevRandao[:],
			BaseFeePerGas:   uint256ToIntToByteSlice(payload.BaseFeePerGas),
			FeeRecipient:    payload.FeeRecipient[:],
			BlockHash:       payload.BlockHash[:],
			ExtraData:       payload.ExtraData,
			BlockNumber:     payload.BlockNumber,
			GasLimit:        payload.GasLimit,
			Timestamp:       payload.Timestamp,
			GasUsed:         payload.GasUsed,
			Transactions:    transactions,
			Withdrawals:     withdrawals,
			BlobGasUsed:     payload.BlobGasUsed,
			ExcessBlobGas:   payload.ExcessBlobGas,
			BlockAccessList: payload.BlockAccessList,
			SlotNumber:      payload.SlotNumber,
		},
		BlobsBundle:       convertFuluBlobBundleToProto(block.BlobsBundle),
		Signature:         block.Signature[:],
		ExecutionRequests: ConvertGloasExecutionRequestToProto(block.ExecutionRequests),
	}
}

// ProtoRequestToGloasRequest converts a SubmitBlockRequest to a Gloas submission. It checks
// only that the parts are there; what they must hold is the relay's to check.
func ProtoRequestToGloasRequest(block *SubmitBlockRequest) (*GloasRequest, error) {
	switch {
	case block.BidTrace == nil:
		return nil, fmt.Errorf("gloas block: bid trace missing")
	case block.ExecutionPayload == nil:
		return nil, fmt.Errorf("gloas block: execution payload missing")
	case block.BlobsBundle == nil:
		return nil, fmt.Errorf("gloas block: blobs bundle missing")
	case block.ExecutionRequests == nil:
		return nil, ErrEmptyExecutionRequests
	case len(block.ParentBlockRoot) != len(phase0.Root{}):
		return nil, fmt.Errorf("gloas block: parent block root must be %d bytes, got %d", len(phase0.Root{}), len(block.ParentBlockRoot))
	}

	value, err := uint256.FromHex(block.BidTrace.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to convert gloas block value %s to uint256: %s", block.BidTrace.Value, err.Error())
	}

	payload := block.ExecutionPayload
	transactions := make([]bellatrix.Transaction, len(payload.Transactions))
	for i, tx := range payload.Transactions {
		transactions[i] = tx.RawData
	}

	withdrawals := make([]*capella.Withdrawal, len(payload.Withdrawals))
	for i, withdrawal := range payload.Withdrawals {
		withdrawals[i] = &capella.Withdrawal{
			ValidatorIndex: phase0.ValidatorIndex(withdrawal.ValidatorIndex),
			Index:          capella.WithdrawalIndex(withdrawal.Index),
			Amount:         phase0.Gwei(withdrawal.Amount),
			Address:        b20(withdrawal.Address),
		}
	}

	blobsBundle := &apiFulu.BlobsBundle{
		Commitments: make([]deneb.KZGCommitment, len(block.BlobsBundle.Commitments)),
		Proofs:      make([]deneb.KZGProof, len(block.BlobsBundle.Proofs)),
		Blobs:       make([]deneb.Blob, len(block.BlobsBundle.Blobs)),
	}
	for i, commitment := range block.BlobsBundle.Commitments {
		copy(blobsBundle.Commitments[i][:], commitment)
	}
	for i, proof := range block.BlobsBundle.Proofs {
		copy(blobsBundle.Proofs[i][:], proof)
	}
	for i, blob := range block.BlobsBundle.Blobs {
		copy(blobsBundle.Blobs[i][:], blob)
	}

	return &GloasRequest{
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
		ParentBlockRoot: b32(block.ParentBlockRoot),
		ExecutionPayload: &gloas.ExecutionPayload{
			ParentHash:      b32(payload.ParentHash),
			FeeRecipient:    b20(payload.FeeRecipient),
			StateRoot:       b32(payload.StateRoot),
			ReceiptsRoot:    b32(payload.ReceiptsRoot),
			LogsBloom:       b256(payload.LogsBloom),
			PrevRandao:      b32(payload.PrevRandao),
			BlockNumber:     payload.BlockNumber,
			GasLimit:        payload.GasLimit,
			GasUsed:         payload.GasUsed,
			Timestamp:       payload.Timestamp,
			ExtraData:       payload.ExtraData,
			BaseFeePerGas:   byteSliceToUint256Int(payload.BaseFeePerGas),
			BlockHash:       b32(payload.BlockHash),
			Transactions:    transactions,
			Withdrawals:     withdrawals,
			BlobGasUsed:     payload.BlobGasUsed,
			ExcessBlobGas:   payload.ExcessBlobGas,
			BlockAccessList: payload.BlockAccessList,
			SlotNumber:      payload.SlotNumber,
		},
		BlobsBundle:       blobsBundle,
		ExecutionRequests: convertGloasProtoToExecutionRequest(block.ExecutionRequests),
		Signature:         b96(block.Signature),
	}, nil
}

// ProtoRequestToGloasBidtracePayload reads what the relay checks before it decodes the
// rest of a Gloas block: the bid trace, the parent block root and the payload's timestamp.
func ProtoRequestToGloasBidtracePayload(block *SubmitBlockRequest) (*BidtracePayload, error) {
	if len(block.ParentBlockRoot) != len(phase0.Root{}) {
		return nil, fmt.Errorf("gloas block: parent block root must be %d bytes, got %d", len(phase0.Root{}), len(block.ParentBlockRoot))
	}
	bidtracePayload, err := ProtoRequestToFuluBidtracePayload(block)
	if err != nil {
		return nil, err
	}
	parentBlockRoot := phase0.Root(b32(block.ParentBlockRoot))
	bidtracePayload.ParentBlockRoot = &parentBlockRoot
	return bidtracePayload, nil
}

// ConvertGloasExecutionRequestToProto converts Gloas execution requests, which add builder
// deposits and exits to Electra's three lists.
func ConvertGloasExecutionRequestToProto(executionRequests *gloas.ExecutionRequests) *ExecutionRequests {
	protoExecutionRequests := ConvertFuluExecutionRequestToProto(&electra.ExecutionRequests{
		Deposits:       executionRequests.Deposits,
		Withdrawals:    executionRequests.Withdrawals,
		Consolidations: executionRequests.Consolidations,
	})

	protoExecutionRequests.BuilderDeposits = make([]*BuilderDepositRequest, len(executionRequests.BuilderDeposits))
	for i, deposit := range executionRequests.BuilderDeposits {
		protoExecutionRequests.BuilderDeposits[i] = &BuilderDepositRequest{
			Pubkey:                deposit.Pubkey[:],
			WithdrawalCredentials: deposit.WithdrawalCredentials,
			Amount:                uint64(deposit.Amount),
			Signature:             deposit.Signature[:],
		}
	}

	protoExecutionRequests.BuilderExits = make([]*BuilderExitRequest, len(executionRequests.BuilderExits))
	for i, exit := range executionRequests.BuilderExits {
		protoExecutionRequests.BuilderExits[i] = &BuilderExitRequest{
			SourceAddress: exit.SourceAddress[:],
			Pubkey:        exit.Pubkey[:],
		}
	}

	return protoExecutionRequests
}

func convertGloasProtoToExecutionRequest(protoExecutionRequests *ExecutionRequests) *gloas.ExecutionRequests {
	electraRequests := convertFuluProtoToExecutionRequest(protoExecutionRequests)
	executionRequests := &gloas.ExecutionRequests{
		Deposits:        electraRequests.Deposits,
		Withdrawals:     electraRequests.Withdrawals,
		Consolidations:  electraRequests.Consolidations,
		BuilderDeposits: make([]*gloas.BuilderDepositRequest, len(protoExecutionRequests.BuilderDeposits)),
		BuilderExits:    make([]*gloas.BuilderExitRequest, len(protoExecutionRequests.BuilderExits)),
	}

	for i, deposit := range protoExecutionRequests.BuilderDeposits {
		executionRequests.BuilderDeposits[i] = &gloas.BuilderDepositRequest{
			Pubkey:                b48(deposit.Pubkey),
			WithdrawalCredentials: deposit.WithdrawalCredentials,
			Amount:                phase0.Gwei(deposit.Amount),
			Signature:             b96(deposit.Signature),
		}
	}

	for i, exit := range protoExecutionRequests.BuilderExits {
		executionRequests.BuilderExits[i] = &gloas.BuilderExitRequest{
			SourceAddress: b20(exit.SourceAddress),
			Pubkey:        b48(exit.Pubkey),
		}
	}

	return executionRequests
}
