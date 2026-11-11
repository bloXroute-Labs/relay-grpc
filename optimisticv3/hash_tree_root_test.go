package optimisticv3

import (
	"encoding/hex"
	"testing"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	"github.com/bloXroute-Labs/relay-grpc/bidadjustment"
)

// The roots below were computed by the code before go-eth2-client v0.30, whose
// types hash with dynamic-ssz's HashWalker rather than fastssz's. Our generated
// code now appends those fields' own roots instead of passing them our hasher;
// these pin that the result is unchanged.
const (
	goldenDenebV2Root = "8db91a34f275ced247b1815245317661b69a8d39efd5941d6b0dcd86e8501133"
	goldenElectraRoot = "00ed7c369c138b6bfbfb38f3f1e5496bcf4a5efb17e44139d24a4b1e2c4c3907"
	goldenFuluRoot    = "00ed7c369c138b6bfbfb38f3f1e5496bcf4a5efb17e44139d24a4b1e2c4c3907"
)

func testHashBidTrace() *v1.BidTrace {
	return &v1.BidTrace{Slot: 12345, GasLimit: 30_000_000, GasUsed: 21_000, Value: uint256.NewInt(27)}
}

func testHashHeader() *deneb.ExecutionPayloadHeader {
	return &deneb.ExecutionPayloadHeader{
		ParentHash:       [32]byte{1},
		FeeRecipient:     bellatrix.ExecutionAddress{2},
		BlockNumber:      100,
		GasLimit:         30_000_000,
		GasUsed:          21_000,
		Timestamp:        12,
		ExtraData:        []byte{3, 4},
		BaseFeePerGas:    uint256.NewInt(7),
		BlockHash:        [32]byte{5},
		TransactionsRoot: [32]byte{6},
		WithdrawalsRoot:  [32]byte{7},
		BlobGasUsed:      131072,
		ExcessBlobGas:    4,
	}
}

func testHashRequests() *electra.ExecutionRequests {
	return &electra.ExecutionRequests{
		Deposits:       []*electra.DepositRequest{{WithdrawalCredentials: make([]byte, 32), Amount: 32, Index: 1}},
		Withdrawals:    []*electra.WithdrawalRequest{{Amount: 2}},
		Consolidations: []*electra.ConsolidationRequest{},
	}
}

func testHashAdjustmentData() bidadjustment.VersionedAdjustmentData {
	return bidadjustment.VersionedAdjustmentData{Version: bidadjustment.AdjustmentDataVersion3, V3: &bidadjustment.AdjustmentDataV3{
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

func TestHeaderSubmissionHashTreeRoot(t *testing.T) {
	commitments := []deneb.KZGCommitment{{18}}

	t.Run("denebV2", func(t *testing.T) {
		h := &HeaderSubmissionDenebV2{BidTrace: testHashBidTrace(), ExecutionPayloadHeader: testHashHeader(), Commitments: commitments}
		requireRoot(t, goldenDenebV2Root, h.HashTreeRoot)
	})
	t.Run("electra", func(t *testing.T) {
		h := &HeaderSubmissionElectra{
			BidTrace:               testHashBidTrace(),
			ExecutionPayloadHeader: testHashHeader(),
			ExecutionRequests:      testHashRequests(),
			Commitments:            commitments,
			AdjustmentData:         testHashAdjustmentData(),
		}
		requireRoot(t, goldenElectraRoot, h.HashTreeRoot)
	})
	t.Run("fulu", func(t *testing.T) {
		h := &HeaderSubmissionFulu{
			BidTrace:               testHashBidTrace(),
			ExecutionPayloadHeader: testHashHeader(),
			ExecutionRequests:      testHashRequests(),
			Commitments:            commitments,
			AdjustmentData:         testHashAdjustmentData(),
		}
		requireRoot(t, goldenFuluRoot, h.HashTreeRoot)
	})
}
