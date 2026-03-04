package bidadjustment

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	builderApiDeneb "github.com/attestantio/go-builder-client/api/deneb"
	builderApiV1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

func TestAdjustableDataSSZ(t *testing.T) {
	filename := "./adjustableSubmitBlockPayloadElectra.ssz"

	sszBytes, err := os.ReadFile(filename)
	require.NoError(t, err)
	req := new(ElectraAdjustableSubmitBlockRequest)

	err = req.UnmarshalSSZ(sszBytes)
	require.NoError(t, err)

	roundtripped, err := req.MarshalSSZ()
	require.NoError(t, err)
	require.True(t, bytes.Equal(roundtripped, sszBytes))
}

func TestAdjustableSubmitBlockRequest(t *testing.T) {
	// Create a new AdjustableSubmitBlockRequest object
	a := &ElectraAdjustableSubmitBlockRequest{
		Signature: phase0.BLSSignature{},
		Message: &builderApiV1.BidTrace{
			ParentHash: phase0.Hash32(common.HexToHash("0x1234567890abcdef")),
			BlockHash:  phase0.Hash32(common.HexToHash("0xabcdef1234567890")),
			GasLimit:   3,
			GasUsed:    4,
			// This value is actual profit + 1, validation should fail
			Value: uint256.NewInt(132912184722469),
		},
		ExecutionPayload: &deneb.ExecutionPayload{
			ParentHash:    phase0.Hash32(common.HexToHash("0xabcdef1234567890")),
			FeeRecipient:  bellatrix.ExecutionAddress{},
			StateRoot:     phase0.Root(common.HexToHash("0x1234567890abcdef")),
			ReceiptsRoot:  phase0.Root(common.HexToHash("0xabcdef1234567890")),
			LogsBloom:     [256]byte{},
			PrevRandao:    [32]byte{},
			BlockNumber:   1,
			GasLimit:      2,
			GasUsed:       3,
			Timestamp:     4,
			ExtraData:     make([]byte, 0),
			BaseFeePerGas: uint256.NewInt(1),
			BlockHash:     phase0.Hash32(common.HexToHash("0xabcdef1234567890")),
			Transactions:  make([]bellatrix.Transaction, 0),
			Withdrawals:   make([]*capella.Withdrawal, 0),
			BlobGasUsed:   5,
			ExcessBlobGas: 6,
		},
		BlobsBundle: &builderApiDeneb.BlobsBundle{
			Commitments: make([]deneb.KZGCommitment, 0),
			Proofs:      make([]deneb.KZGProof, 0),
			Blobs:       make([]deneb.Blob, 0),
		},
		AdjustmentData: &VersionedAdjustmentData{
			Version: AdjustmentDataVersion1,
			V1: &AdjustmentData{
				StateRoot:                   phase0.Hash32(common.HexToHash("0xabcdef1234567890")),
				TransactionsRoot:            phase0.Hash32(common.HexToHash("0x1234567890abcdef")),
				ReceiptsRoot:                phase0.Hash32(common.HexToHash("0xabcdef1234567890")),
				BuilderAddress:              common.HexToAddress("0x1234567890abcdef"),
				BuilderProof:                make([][]byte, 0),
				FeeRecipientAddress:         common.HexToAddress("0xabcdef1234567890"),
				FeeRecipientProof:           make([][]byte, 0),
				FeePayerAddress:             common.HexToAddress("0x1234567890abcdef"),
				FeePayerProof:               make([][]byte, 0),
				PlaceholderTransactionProof: make([][]byte, 0),
				PlaceholderReceiptProof:     make([][]byte, 0),
			},
		},
		ExecutionRequests: &electra.ExecutionRequests{
			Deposits:       make([]*electra.DepositRequest, 0),
			Withdrawals:    make([]*electra.WithdrawalRequest, 0),
			Consolidations: make([]*electra.ConsolidationRequest, 0),
		},
	}
	sszA, err := a.MarshalSSZ()
	if err != nil {
		t.Fatalf("Failed to marshal SSZ: %v", err)
	}
	newA := &ElectraAdjustableSubmitBlockRequest{}
	if err := newA.UnmarshalSSZ(sszA); err != nil {
		t.Fatalf("Failed to unmarshal SSZ: %v", err)
	}
	if a.Signature != newA.Signature {
		t.Fatalf("Signature mismatch")
	}
	if a.Message.ParentHash != newA.Message.ParentHash {
		t.Fatalf("Message.ParentHash mismatch")
	}
	if a.Message.BlockHash != newA.Message.BlockHash {
		t.Fatalf("Message.BlockHash mismatch")
	}
	if a.Message.GasLimit != newA.Message.GasLimit {
		t.Fatalf("Message.GasLimit mismatch")
	}
	if a.Message.GasUsed != newA.Message.GasUsed {
		t.Fatalf("Message.GasUsed mismatch")
	}
	if a.Message.Value.Cmp(newA.Message.Value) != 0 {
		t.Fatalf("Message.Value mismatch")
	}
	if a.ExecutionPayload.ParentHash != newA.ExecutionPayload.ParentHash {
		t.Fatalf("ExecutionPayload.ParentHash mismatch")
	}
	if len(a.BlobsBundle.Commitments) != len(newA.BlobsBundle.Commitments) {
		t.Fatalf("BlobsBundle.Commitments mismatch")
	}
	if len(a.BlobsBundle.Proofs) != len(newA.BlobsBundle.Proofs) {
		t.Fatalf("BlobsBundle.Proofs mismatch")
	}
	if len(a.BlobsBundle.Blobs) != len(newA.BlobsBundle.Blobs) {
		t.Fatalf("BlobsBundle.Blobs mismatch")
	}
	if a.AdjustmentData.V1.StateRoot != newA.AdjustmentData.V1.StateRoot {
		t.Fatalf("AdjustmentData.StateRoot mismatch")
	}
	if a.AdjustmentData.V1.TransactionsRoot != newA.AdjustmentData.V1.TransactionsRoot {
		t.Fatalf("AdjustmentData.TransactionsRoot mismatch")
	}
	if a.AdjustmentData.V1.ReceiptsRoot != newA.AdjustmentData.V1.ReceiptsRoot {
		t.Fatalf("AdjustmentData.ReceiptsRoot mismatch")
	}
	if a.AdjustmentData.V1.BuilderAddress != newA.AdjustmentData.V1.BuilderAddress {
		t.Fatalf("AdjustmentData.BuilderAddress mismatch")
	}
	if len(a.AdjustmentData.V1.BuilderProof) != len(newA.AdjustmentData.V1.BuilderProof) {
		t.Fatalf("AdjustmentData.BuilderProof mismatch")
	}
	if a.AdjustmentData.V1.FeeRecipientAddress != newA.AdjustmentData.V1.FeeRecipientAddress {
		t.Fatalf("AdjustmentData.FeeRecipientAddress mismatch")
	}
	if len(a.AdjustmentData.V1.FeeRecipientProof) != len(newA.AdjustmentData.V1.FeeRecipientProof) {
		t.Fatalf("AdjustmentData.FeeRecipientProof mismatch")
	}
	if a.AdjustmentData.V1.FeePayerAddress != newA.AdjustmentData.V1.FeePayerAddress {
		t.Fatalf("AdjustmentData.FeePayerAddress mismatch")
	}
	if len(a.AdjustmentData.V1.FeePayerProof) != len(newA.AdjustmentData.V1.FeePayerProof) {
		t.Fatalf("AdjustmentData.FeePayerProof mismatch")
	}
	if len(a.AdjustmentData.V1.PlaceholderTransactionProof) != len(newA.AdjustmentData.V1.PlaceholderTransactionProof) {
		t.Fatalf("AdjustmentData.PlaceholderTransactionProof mismatch")
	}
	if len(a.AdjustmentData.V1.PlaceholderReceiptProof) != len(newA.AdjustmentData.V1.PlaceholderReceiptProof) {
		t.Fatalf("AdjustmentData.PlaceholderReceiptProof mismatch")
	}
}
