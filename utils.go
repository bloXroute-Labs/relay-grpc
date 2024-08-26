package relay_grpc

import (
	"fmt"
	"math/big"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	builderSpec "github.com/attestantio/go-builder-client/spec"
	consensusspec "github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
)

var ErrInvalidVersion = errors.New("invalid version")

// Based on the version, delegate to the correct RequestToProtoRequest
func VersionedRequestToProtoRequest(block *builderSpec.VersionedSubmitBlockRequest) (*SubmitBlockRequest, error) {
	switch block.Version {
	case consensusspec.DataVersionCapella:
		return CapellaRequestToProtoRequest(block.Capella), nil
	case consensusspec.DataVersionDeneb:
		return DenebRequestToProtoRequest(block.Deneb), nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%s is not supported", block.Version))
	}
}

// Based on the version, delegate to the correct RequestToProtoRequestWithShortIDs
func VersionedRequestToProtoRequestWithShortIDs(block *builderSpec.VersionedSubmitBlockRequest, compressTxs []*CompressTx) (*SubmitBlockRequest, error) {
	switch block.Version {
	case consensusspec.DataVersionCapella:
		return CapellaRequestToProtoRequestWithShortIDs(block.Capella, compressTxs), nil
	case consensusspec.DataVersionDeneb:
		return DenebRequestToProtoRequestWithShortIDs(block.Deneb, compressTxs), nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%s is not supported", block.Version))
	}
}

// Based on the version, delegate to the correct ProtoRequestToVersionedRequest
func ProtoRequestToVersionedRequest(block *SubmitBlockRequest) (*builderSpec.VersionedSubmitBlockRequest, error) {
	switch consensusspec.DataVersion(block.Version) {
	case consensusspec.DataVersionCapella:
		blockRequest, err := ProtoRequestToCapellaRequest(block)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSubmitBlockRequest{
			Version: consensusspec.DataVersionCapella,
			Capella: blockRequest,
		}, nil
	case consensusspec.DataVersionDeneb:
		blockRequest, err := ProtoRequestToDenebRequest(block)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSubmitBlockRequest{
			Version: consensusspec.DataVersionDeneb,
			Deneb:   blockRequest,
		}, nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%s is not supported", consensusspec.DataVersion(block.Version)))
	}
}

type BidTraceExecutionPayload struct {
	Timestamp uint64
}
type BidtracePayload struct {
	Message          *v1.BidTrace
	ExecutionPayload *BidTraceExecutionPayload
	Signature        phase0.BLSSignature
}

func ProtoRequestToBidtracePayload(block *SubmitBlockRequest) (*BidtracePayload, error) {
	blockRequest, err := ProtoRequestToDenebBidtracePayload(block)
	if err != nil {
		return nil, err
	}
	return blockRequest, nil
}

// b20 converts a byte slice to a [20]byte.
func b20(b []byte) [20]byte {
	out := [20]byte{}
	copy(out[:], b)
	return out
}

// b32 converts a byte slice to a [32]byte.
func b32(b []byte) [32]byte {
	out := [32]byte{}
	copy(out[:], b)
	return out
}

// b48 converts a byte slice to a [48]byte.
func b48(b []byte) [48]byte {
	out := [48]byte{}
	copy(out[:], b)
	return out
}

// b96 converts a byte slice to a [96]byte.
func b96(b []byte) [96]byte {
	out := [96]byte{}
	copy(out[:], b)
	return out
}

// b256 converts a byte slice to a [256]byte.
func b256(b []byte) [256]byte {
	out := [256]byte{}
	copy(out[:], b)
	return out
}

// uint256ToIntToByteSlice converts a *uint256.Int to a byte slice.
func uint256ToIntToByteSlice(u *uint256.Int) []byte {
	if u == nil {
		return nil
	}
	// Convert the uint256.Int to a byte slice.
	// The Bytes method returns the absolute value as a big-endian byte slice.
	return u.Bytes()
}

// byteSliceToUint256Int converts a byte slice to a *uint256.Int.
func byteSliceToUint256Int(b []byte) *uint256.Int {
	u256, _ := uint256.FromBig(new(big.Int).SetBytes(b))
	return u256
}
