package relay_grpc

import (
	"github.com/attestantio/go-builder-client/api/capella"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	consensus "github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

func CapellaRequestToProtoRequest(block *capella.SubmitBlockRequest) *SubmitBlockRequest {

	transactions := []*CompressTx{}
	for _, tx := range block.ExecutionPayload.Transactions {
		transactions = append(transactions, &CompressTx{
			RawData: []byte(tx),
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
		},
		ExecutionPayload: &ExecutionPayload{
			ParentHash:    block.ExecutionPayload.ParentHash[:],
			StateRoot:     block.ExecutionPayload.StateRoot[:],
			ReceiptsRoot:  block.ExecutionPayload.ReceiptsRoot[:],
			LogsBloom:     block.ExecutionPayload.LogsBloom[:],
			PrevRandao:    block.ExecutionPayload.PrevRandao[:],
			BaseFeePerGas: block.ExecutionPayload.BaseFeePerGas[:],
			FeeRecipient:  block.ExecutionPayload.FeeRecipient[:],
			BlockHash:     block.ExecutionPayload.BlockHash[:],
			ExtraData:     block.ExecutionPayload.ExtraData,
			BlockNumber:   block.ExecutionPayload.BlockNumber,
			GasLimit:      block.ExecutionPayload.GasLimit,
			Timestamp:     block.ExecutionPayload.Timestamp,
			GasUsed:       block.ExecutionPayload.GasUsed,
			Transactions:  transactions,
			Withdrawals:   withdrawals,
		},
		Signature: block.Signature[:],
	}
}

func CapellaRequestToProtoRequestWithShortIDs(block *capella.SubmitBlockRequest, compressTxs []*CompressTx) *SubmitBlockRequest {

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
		},
		ExecutionPayload: &ExecutionPayload{
			ParentHash:    block.ExecutionPayload.ParentHash[:],
			StateRoot:     block.ExecutionPayload.StateRoot[:],
			ReceiptsRoot:  block.ExecutionPayload.ReceiptsRoot[:],
			LogsBloom:     block.ExecutionPayload.LogsBloom[:],
			PrevRandao:    block.ExecutionPayload.PrevRandao[:],
			BaseFeePerGas: block.ExecutionPayload.BaseFeePerGas[:],
			FeeRecipient:  block.ExecutionPayload.FeeRecipient[:],
			BlockHash:     block.ExecutionPayload.BlockHash[:],
			ExtraData:     block.ExecutionPayload.ExtraData,
			BlockNumber:   block.ExecutionPayload.BlockNumber,
			GasLimit:      block.ExecutionPayload.GasLimit,
			Timestamp:     block.ExecutionPayload.Timestamp,
			GasUsed:       block.ExecutionPayload.GasUsed,
			Transactions:  compressTxs,
			Withdrawals:   withdrawals,
		},
		Signature: block.Signature[:],
	}
}

func b20(b []byte) [20]byte {
	out := [20]byte{}
	copy(out[:], b)
	return out
}
func b32(b []byte) [32]byte {
	out := [32]byte{}
	copy(out[:], b)
	return out
}
func b48(b []byte) [48]byte {
	out := [48]byte{}
	copy(out[:], b)
	return out
}
func b96(b []byte) [96]byte {
	out := [96]byte{}
	copy(out[:], b)
	return out
}
func b256(b []byte) [256]byte {
	out := [256]byte{}
	copy(out[:], b)
	return out
}

func ProtoRequestToCapellaRequest(block *SubmitBlockRequest) *capella.SubmitBlockRequest {

	transactions := []bellatrix.Transaction{}
	for _, tx := range block.ExecutionPayload.Transactions {
		transactions = append(transactions, tx.RawData)
	}

	withdrawals := []*consensus.Withdrawal{}

	for _, withdrawal := range block.ExecutionPayload.Withdrawals {
		withdrawals = append(withdrawals, &consensus.Withdrawal{
			ValidatorIndex: phase0.ValidatorIndex(withdrawal.ValidatorIndex),
			Index:          consensus.WithdrawalIndex(withdrawal.Index),
			Amount:         phase0.Gwei(withdrawal.Amount),
			Address:        b20(withdrawal.Address),
		})
	}

	return &capella.SubmitBlockRequest{
		Message: &v1.BidTrace{
			Slot:                 block.BidTrace.Slot,
			ParentHash:           b32(block.BidTrace.ParentHash),
			BlockHash:            b32(block.BidTrace.BlockHash),
			BuilderPubkey:        b48(block.BidTrace.BuilderPubkey),
			ProposerPubkey:       b48(block.BidTrace.ProposerPubkey),
			ProposerFeeRecipient: b20(block.BidTrace.ProposerFeeRecipient),
			GasLimit:             block.BidTrace.GasLimit,
			GasUsed:              block.BidTrace.GasUsed,
			Value:                uint256.MustFromHex(block.BidTrace.Value),
		},
		ExecutionPayload: &consensus.ExecutionPayload{
			ParentHash:    b32(block.ExecutionPayload.ParentHash),
			StateRoot:     b32(block.ExecutionPayload.StateRoot),
			ReceiptsRoot:  b32(block.ExecutionPayload.ReceiptsRoot),
			LogsBloom:     b256(block.ExecutionPayload.LogsBloom),
			PrevRandao:    b32(block.ExecutionPayload.PrevRandao),
			BaseFeePerGas: b32(block.ExecutionPayload.BaseFeePerGas),
			FeeRecipient:  b20(block.ExecutionPayload.FeeRecipient),
			BlockHash:     b32(block.ExecutionPayload.BlockHash),
			ExtraData:     block.ExecutionPayload.ExtraData,
			BlockNumber:   block.ExecutionPayload.BlockNumber,
			GasLimit:      block.ExecutionPayload.GasLimit,
			Timestamp:     block.ExecutionPayload.Timestamp,
			GasUsed:       block.ExecutionPayload.GasUsed,
			Transactions:  transactions,
			Withdrawals:   withdrawals,
		},
		Signature: b96(block.Signature),
	}
}
