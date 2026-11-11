package relay_grpc

import (
	"testing"

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
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func testGloasRequest() *GloasRequest {
	return &GloasRequest{
		Message: &v1.BidTrace{
			Slot:                 64,
			ParentHash:           phase0.Hash32{0x01},
			BlockHash:            phase0.Hash32{0x02},
			BuilderPubkey:        phase0.BLSPubKey{0x03},
			ProposerPubkey:       phase0.BLSPubKey{0x04},
			ProposerFeeRecipient: bellatrix.ExecutionAddress{0x05},
			GasLimit:             60_000_000,
			GasUsed:              21_000,
			Value:                uint256.NewInt(1_000_000_000),
		},
		ParentBlockRoot: phase0.Root{0x21},
		ExecutionPayload: &gloas.ExecutionPayload{
			ParentHash:      phase0.Hash32{0x01},
			FeeRecipient:    bellatrix.ExecutionAddress{0x06},
			StateRoot:       phase0.Root{0x07},
			ReceiptsRoot:    phase0.Root{0x08},
			LogsBloom:       [256]byte{0x09},
			PrevRandao:      [32]byte{0x0a},
			BlockNumber:     10,
			GasLimit:        60_000_000,
			GasUsed:         21_000,
			Timestamp:       1_700_000_000,
			ExtraData:       []byte("blxr"),
			BaseFeePerGas:   uint256.NewInt(7),
			BlockHash:       phase0.Hash32{0x02},
			Transactions:    []bellatrix.Transaction{{0x02, 0x01}},
			Withdrawals:     []*capella.Withdrawal{{Index: 1, ValidatorIndex: 2, Address: bellatrix.ExecutionAddress{0x0b}, Amount: 3}},
			BlobGasUsed:     131072,
			ExcessBlobGas:   393216,
			BlockAccessList: gloas.BlockAccessList{0xc1, 0xc0},
			SlotNumber:      64,
		},
		BlobsBundle: &apiFulu.BlobsBundle{
			Commitments: []deneb.KZGCommitment{{0x0c}},
			Proofs:      make([]deneb.KZGProof, 128),
			Blobs:       []deneb.Blob{{0x0d}},
		},
		ExecutionRequests: &gloas.ExecutionRequests{
			Deposits:        []*electra.DepositRequest{{Pubkey: phase0.BLSPubKey{0x0e}, WithdrawalCredentials: make([]byte, 32), Amount: 1, Index: 2}},
			Withdrawals:     []*electra.WithdrawalRequest{{SourceAddress: bellatrix.ExecutionAddress{0x0f}, Amount: 3}},
			Consolidations:  []*electra.ConsolidationRequest{{SourceAddress: bellatrix.ExecutionAddress{0x10}}},
			BuilderDeposits: []*gloas.BuilderDepositRequest{{Pubkey: phase0.BLSPubKey{0x11}, WithdrawalCredentials: make([]byte, 32), Amount: 4, Signature: phase0.BLSSignature{0x12}}},
			BuilderExits:    []*gloas.BuilderExitRequest{{SourceAddress: bellatrix.ExecutionAddress{0x13}, Pubkey: phase0.BLSPubKey{0x14}}},
		},
		Signature: phase0.BLSSignature{0x15},
	}
}

// A Gloas submission survives the trip through the proto and its wire encoding, including
// what Gloas adds: the parent block root, the block access list, the slot number and the
// builder deposits and exits.
func TestGloasRequestProtoRoundTrip(t *testing.T) {
	request := testGloasRequest()

	protoRequest := GloasRequestToProtoRequest(request)
	require.Equal(t, uint64(consensusspec.DataVersionGloas), protoRequest.Version)
	wire, err := proto.Marshal(protoRequest)
	require.NoError(t, err)
	decoded := new(SubmitBlockRequest)
	require.NoError(t, proto.Unmarshal(wire, decoded))

	got, err := ProtoRequestToGloasRequest(decoded)
	require.NoError(t, err)
	require.Equal(t, request, got)

	bidtracePayload, err := ProtoRequestToBidtracePayload(decoded)
	require.NoError(t, err)
	require.Equal(t, request.Message, bidtracePayload.Message)
	require.Equal(t, request.ExecutionPayload.Timestamp, bidtracePayload.ExecutionPayload.Timestamp)
	require.Equal(t, request.ParentBlockRoot, *bidtracePayload.ParentBlockRoot)
}

// A Gloas block without its parent block root, or with a malformed one, is refused rather
// than read as the zero root.
func TestProtoRequestToGloasRequestNeedsTheParentBlockRoot(t *testing.T) {
	for name, root := range map[string][]byte{"missing": nil, "short": {0x21}} {
		t.Run(name, func(t *testing.T) {
			protoRequest := GloasRequestToProtoRequest(testGloasRequest())
			protoRequest.ParentBlockRoot = root
			_, err := ProtoRequestToGloasRequest(protoRequest)
			require.ErrorContains(t, err, "parent block root")
			_, err = ProtoRequestToBidtracePayload(protoRequest)
			require.ErrorContains(t, err, "parent block root")
		})
	}
}
