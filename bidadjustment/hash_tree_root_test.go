package bidadjustment

import (
	"encoding/hex"
	"testing"

	d "github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-builder-client/api/fulu"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

// The roots below were computed by the code before go-eth2-client v0.30, whose
// types hash with dynamic-ssz's HashWalker rather than fastssz's. Our generated
// code now appends those fields' own roots instead of passing them our hasher;
// these pin that the result is unchanged.
const (
	goldenDenebRoot   = "b97f0f30e90208a30fa474dd1b73cf05e4660d5020694b2f099051555ef61063"
	goldenElectraRoot = "bdef6cabd47092037b5feaad2150f8a5ed7055175858e8c729dc42cdb6bbe794"
	goldenFuluRoot    = "55fd54baecdf454be5b9d16fc2e0e3ff2869743c6b09631cdbd33ad066b7bc6d"
)

func testHashMessage() *v1.BidTrace {
	return &v1.BidTrace{Slot: 12345, GasLimit: 30_000_000, GasUsed: 21_000, Value: uint256.NewInt(27)}
}

func testHashPayload() *deneb.ExecutionPayload {
	return &deneb.ExecutionPayload{
		ParentHash:    [32]byte{1},
		FeeRecipient:  bellatrix.ExecutionAddress{2},
		BlockNumber:   100,
		GasLimit:      30_000_000,
		GasUsed:       21_000,
		Timestamp:     12,
		ExtraData:     []byte{3, 4},
		BaseFeePerGas: uint256.NewInt(7),
		BlockHash:     [32]byte{5},
		Transactions:  []bellatrix.Transaction{{6, 7, 8}, {9}},
		Withdrawals:   []*capella.Withdrawal{{Index: 1, ValidatorIndex: 2, Address: bellatrix.ExecutionAddress{10}, Amount: 3}},
		BlobGasUsed:   131072,
		ExcessBlobGas: 4,
	}
}

func testHashRequests() *electra.ExecutionRequests {
	return &electra.ExecutionRequests{
		Deposits:       []*electra.DepositRequest{{WithdrawalCredentials: make([]byte, 32), Amount: 32, Index: 1}},
		Withdrawals:    []*electra.WithdrawalRequest{{Amount: 2}},
		Consolidations: []*electra.ConsolidationRequest{},
	}
}

func testHashAdjustmentData() *VersionedAdjustmentData {
	return &VersionedAdjustmentData{Version: AdjustmentDataVersion3, V3: &AdjustmentDataV3{
		BuilderProof:                  [][]byte{{11}},
		FeeRecipientProof:             [][]byte{{12}},
		FeePayerProof:                 [][]byte{{13}},
		ELPlaceholderTransactionProof: [][]byte{{14}},
		CLPlaceholderTransactionProof: [][32]byte{{15}},
		ELPlaceholderReceiptProof:     [][]byte{{16}},
		PlaceholderGasUsed:            21_000,
	}}
}

func requireRoot(t *testing.T, golden string, hashTreeRoot func() ([32]byte, error)) {
	t.Helper()
	root, err := hashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, golden, hex.EncodeToString(root[:]))
}

func TestAdjustableSubmitBlockRequestHashTreeRoot(t *testing.T) {
	blob := deneb.Blob{17}

	t.Run("deneb", func(t *testing.T) {
		r := &DenebAdjustableSubmitBlockRequest{
			Message:          testHashMessage(),
			ExecutionPayload: testHashPayload(),
			BlobsBundle:      &d.BlobsBundle{Commitments: []deneb.KZGCommitment{{18}}, Proofs: []deneb.KZGProof{{19}}, Blobs: []deneb.Blob{blob}},
			AdjustmentData:   testHashAdjustmentData(),
		}
		requireRoot(t, goldenDenebRoot, r.HashTreeRoot)
	})
	t.Run("electra", func(t *testing.T) {
		r := &ElectraAdjustableSubmitBlockRequest{
			Message:           testHashMessage(),
			ExecutionPayload:  testHashPayload(),
			BlobsBundle:       &d.BlobsBundle{Commitments: []deneb.KZGCommitment{{18}}, Proofs: []deneb.KZGProof{{19}}, Blobs: []deneb.Blob{blob}},
			ExecutionRequests: testHashRequests(),
			AdjustmentData:    testHashAdjustmentData(),
		}
		requireRoot(t, goldenElectraRoot, r.HashTreeRoot)
	})
	t.Run("fulu", func(t *testing.T) {
		r := &FuluAdjustableSubmitBlockRequest{
			Message:           testHashMessage(),
			ExecutionPayload:  testHashPayload(),
			BlobsBundle:       &fulu.BlobsBundle{Commitments: []deneb.KZGCommitment{{18}}, Proofs: []deneb.KZGProof{{19}, {20}}, Blobs: []deneb.Blob{blob}},
			ExecutionRequests: testHashRequests(),
			AdjustmentData:    testHashAdjustmentData(),
		}
		requireRoot(t, goldenFuluRoot, r.HashTreeRoot)
	})
}
