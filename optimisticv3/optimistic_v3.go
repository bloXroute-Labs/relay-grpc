package optimisticv3

import (
	"encoding/json"
	"fmt"

	builderApiDeneb "github.com/attestantio/go-builder-client/api/deneb"
	builderApiElectra "github.com/attestantio/go-builder-client/api/electra"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	builderSpec "github.com/attestantio/go-builder-client/spec"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	relaygrpc "github.com/bloXroute-Labs/relay-grpc"
	"github.com/bloXroute-Labs/relay-grpc/bidadjustment"
	"github.com/flashbots/go-boost-utils/bls"
	"github.com/flashbots/go-boost-utils/ssz"
	"github.com/pkg/errors"
)

const (
	PayloadURLMaxLength           = 256
	PayloadUrlsDataExpectedLength = 2
	PayloadUrlTypeIndex           = 0
	PayloadUrlsCSVIndex           = 1
	PayloadUrlsTypeSeparator      = ";"
)

type PayloadUrlType string

const (
	PayloadUrlTypeGRPC PayloadUrlType = "grpc"
	PayloadUrlTypeHTTP PayloadUrlType = "http"
)

var (
	ErrInvalidVersion   = errors.New("invalid version")
	errMissingRequest   = errors.New("req is nil")
	ErrEmptyPayload     = errors.New("empty payload")
	errMissingSecretKey = errors.New("secret key is nil")
)

func (h *VersionedSignedHeaderSubmission) MarshalJSON() ([]byte, error) {
	switch h.Version { //nolint:exhaustive
	case spec.DataVersionFulu:
		return json.Marshal(h.Fulu)
	case spec.DataVersionElectra:
		return json.Marshal(h.Electra)
	case spec.DataVersionDeneb:
		return json.Marshal(h.Deneb)
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%s is not supported", h.Version))
	}
}

func (h *VersionedSignedHeaderSubmission) UnmarshalJSON(input []byte) error {
	var err error

	fuluRequest := new(SignedHeaderSubmissionFulu)
	if err = json.Unmarshal(input, fuluRequest); err == nil {
		h.Version = spec.DataVersionFulu
		h.Fulu = fuluRequest
		return nil
	}

	electraRequest := new(SignedHeaderSubmissionElectra)
	if err = json.Unmarshal(input, electraRequest); err == nil {
		h.Version = spec.DataVersionElectra
		h.Electra = electraRequest
		return nil
	}

	denebRequest := new(SignedHeaderSubmissionDeneb)
	if err = json.Unmarshal(input, denebRequest); err == nil {
		h.Version = spec.DataVersionDeneb
		h.Deneb = denebRequest
		return nil
	}

	return errors.Wrap(err, "failed to unmarshal SubmitBlockRequest ")
}

func (h *VersionedSignedHeaderSubmission) BidTrace() (*v1.BidTrace, error) {
	if h == nil {
		return nil, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return nil, errors.New("no data")
		}
		return h.Fulu.Message.BidTrace, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return nil, errors.New("no data")
		}
		return h.Electra.Message.BidTrace, nil
	case spec.DataVersionDeneb:
		if h.Deneb == nil {
			return nil, errors.New("no data")
		}
		return h.Deneb.Message.BidTrace, nil
	default:
		return nil, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) ExecutionPayloadHeader() (*deneb.ExecutionPayloadHeader, error) {
	if h == nil {
		return nil, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return nil, errors.New("no data")
		}
		return h.Fulu.Message.ExecutionPayloadHeader, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return nil, errors.New("no data")
		}
		return h.Electra.Message.ExecutionPayloadHeader, nil
	case spec.DataVersionDeneb:
		if h.Deneb == nil {
			return nil, errors.New("no data")
		}
		return h.Deneb.Message.ExecutionPayloadHeader, nil
	default:
		return nil, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) Commitments() ([]deneb.KZGCommitment, error) {
	if h == nil {
		return nil, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return nil, errors.New("no data")
		}
		return h.Fulu.Message.Commitments, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return nil, errors.New("no data")
		}
		return h.Electra.Message.Commitments, nil
	case spec.DataVersionDeneb:
		if h.Deneb == nil {
			return nil, errors.New("no data")
		}
		return h.Deneb.Message.Commitments, nil
	default:
		return nil, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) ExecutionRequests() (*electra.ExecutionRequests, error) {
	if h == nil {
		return nil, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return nil, errors.New("no data")
		}
		return h.Fulu.Message.ExecutionRequests, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return nil, errors.New("no data")
		}
		return h.Electra.Message.ExecutionRequests, nil
	case spec.DataVersionDeneb:
		return nil, errors.New("no executionRequests in deneb")
	default:
		return nil, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) Signature() (phase0.BLSSignature, error) {
	if h == nil {
		return phase0.BLSSignature{}, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return phase0.BLSSignature{}, errors.New("no data")
		}
		return h.Fulu.Signature, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return phase0.BLSSignature{}, errors.New("no data")
		}
		return h.Electra.Signature, nil
	case spec.DataVersionDeneb:
		if h.Deneb == nil {
			return phase0.BLSSignature{}, errors.New("no data")
		}
		return h.Deneb.Signature, nil
	default:
		return phase0.BLSSignature{}, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) TxRoot() (phase0.Root, error) {
	if h == nil {
		return phase0.Root{}, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return phase0.Root{}, errors.New("no data")
		}
		return h.Fulu.Message.ExecutionPayloadHeader.TransactionsRoot, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return phase0.Root{}, errors.New("no data")
		}
		return h.Electra.Message.ExecutionPayloadHeader.TransactionsRoot, nil
	case spec.DataVersionDeneb:
		if h.Deneb == nil {
			return phase0.Root{}, errors.New("no data")
		}
		return h.Deneb.Message.ExecutionPayloadHeader.TransactionsRoot, nil
	default:
		return phase0.Root{}, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) WithdrawalsRoot() (phase0.Root, error) {
	if h == nil {
		return phase0.Root{}, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return phase0.Root{}, errors.New("no data")
		}
		return h.Fulu.Message.ExecutionPayloadHeader.WithdrawalsRoot, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return phase0.Root{}, errors.New("no data")
		}
		return h.Electra.Message.ExecutionPayloadHeader.WithdrawalsRoot, nil
	case spec.DataVersionDeneb:
		if h.Deneb == nil {
			return phase0.Root{}, errors.New("no data")
		}
		return h.Deneb.Message.ExecutionPayloadHeader.WithdrawalsRoot, nil
	default:
		return phase0.Root{}, errors.New("unsupported version")
	}
}

func (h *VersionedSignedHeaderSubmission) AdjustmentData() (*bidadjustment.AdjustmentDataV2, error) {
	if h == nil {
		return nil, errors.New("nil struct")
	}
	switch h.Version {
	case spec.DataVersionFulu:
		if h.Fulu == nil {
			return nil, errors.New("no data")
		}
		return &h.Fulu.Message.AdjustmentData, nil
	case spec.DataVersionElectra:
		if h.Electra == nil {
			return nil, errors.New("no data")
		}
		return &h.Electra.Message.AdjustmentData, nil
	case spec.DataVersionDeneb:
		return nil, errors.New("no data")
	default:
		return nil, errors.New("unsupported version")
	}
}

func BuildGetHeaderResponseV3(payload *HeaderSubmissionV3, sk *bls.SecretKey, pubkey *phase0.BLSPubKey, domain phase0.Domain) (*builderSpec.VersionedSignedBuilderBid, error) {
	if payload == nil {
		return nil, errMissingRequest
	}

	if sk == nil {
		return nil, errMissingSecretKey
	}

	switch payload.Submission.Version {
	case spec.DataVersionFulu:
		signedBuilderBid, err := BuilderBlockRequestToSignedBuilderBidV3(payload, sk, pubkey, domain)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSignedBuilderBid{
			Version: spec.DataVersionFulu,
			Fulu:    signedBuilderBid.Fulu,
		}, nil

	case spec.DataVersionElectra:
		signedBuilderBid, err := BuilderBlockRequestToSignedBuilderBidV3(payload, sk, pubkey, domain)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSignedBuilderBid{
			Version: spec.DataVersionElectra,
			Electra: signedBuilderBid.Electra,
		}, nil

	case spec.DataVersionDeneb:
		signedBuilderBid, err := BuilderBlockRequestToSignedBuilderBidV3(payload, sk, pubkey, domain)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSignedBuilderBid{
			Version: spec.DataVersionDeneb,
			Deneb:   signedBuilderBid.Deneb,
		}, nil

	case spec.DataVersionUnknown, spec.DataVersionPhase0, spec.DataVersionAltair, spec.DataVersionBellatrix:
		return nil, ErrInvalidVersion
	default:
		return nil, ErrEmptyPayload
	}
}

type HeaderSubmissionV3 struct {
	// URL pointing to the builder's server endpoint for retrieving
	// the full block payload if this header is selected.
	URL []byte `json:"url" ssz-max:"256"`
	// The number of transactions in the block
	TxCount uint32 `json:"tx_count"`
	// The signed header data. Carrying: ExecutionHeader, BidTrace, Signature
	Submission *VersionedSignedHeaderSubmission `json:"submission"`
}

type VersionedSignedHeaderSubmission struct {
	Version spec.DataVersion
	Deneb   *SignedHeaderSubmissionDeneb   `json:"deneb,omitempty"`
	Electra *SignedHeaderSubmissionElectra `json:"electra,omitempty"`
	Fulu    *SignedHeaderSubmissionFulu    `json:"fulu,omitempty"`
}

type SignedHeaderSubmissionDeneb struct {
	Message   HeaderSubmissionDenebV2 `json:"message"`
	Signature phase0.BLSSignature     `json:"signature" ssz-size:"96"`
}

type SignedHeaderSubmissionElectra struct {
	Message   HeaderSubmissionElectra `json:"message"`
	Signature phase0.BLSSignature     `json:"signature" ssz-size:"96"`
}

type SignedHeaderSubmissionFulu struct {
	Message   HeaderSubmissionFulu `json:"message"`
	Signature phase0.BLSSignature  `json:"signature" ssz-size:"96"`
}

type HeaderSubmissionDenebV2 struct {
	BidTrace               *v1.BidTrace                  `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader `json:"execution_payload_header"`
	Commitments            []deneb.KZGCommitment         `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
}

type HeaderSubmissionElectra struct {
	BidTrace               *v1.BidTrace                   `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader  `json:"execution_payload_header"`
	ExecutionRequests      *electra.ExecutionRequests     `json:"execution_requests"`
	Commitments            []deneb.KZGCommitment          `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
	AdjustmentData         bidadjustment.AdjustmentDataV2 `json:"adjustment_data"`
}

type HeaderSubmissionFulu struct {
	BidTrace               *v1.BidTrace                   `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader  `json:"execution_payload_header"`
	ExecutionRequests      *electra.ExecutionRequests     `json:"execution_requests"`
	Commitments            []deneb.KZGCommitment          `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
	AdjustmentData         bidadjustment.AdjustmentDataV2 `json:"adjustment_data"`
}

type GetPayloadV3 struct {
	// Hash of the block header from the `SignedHeaderSubmission`.
	BlockHash phase0.Hash32 `json:"block_hash" ssz-size:"32"`
	// Timestamp (in milliseconds) when the relay made this request.
	RequestTs uint64 `json:"request_ts"`
	// Bls public key of the signing key that was used to create the `signature` field in `SignedGetPayloadV3`.
	RelayPublicKey phase0.BLSPubKey `json:"relay_public_key" ssz-size:"48"`
}

type SignedGetPayloadV3 struct {
	Message *GetPayloadV3 `json:"message"`
	// Signature from the relay's key that it uses to sign the `get_header` responses.
	Signature phase0.BLSSignature `json:"signature" ssz-size:"96"`
}

type VersionedAdjustableSubmitBlockRequest struct {
	Version spec.DataVersion
	Deneb   *bidadjustment.DenebAdjustableSubmitBlockRequest
	Electra *bidadjustment.ElectraAdjustableSubmitBlockRequest
	Fulu    *bidadjustment.FuluAdjustableSubmitBlockRequest
}

func (v *VersionedAdjustableSubmitBlockRequest) UnmarshalSSZForFork(data []byte, fork spec.DataVersion) error {
	switch fork {
	case spec.DataVersionFulu:
		req := new(bidadjustment.FuluAdjustableSubmitBlockRequest)
		if err := req.UnmarshalSSZ(data); err != nil {
			return err
		}
		v.Version = spec.DataVersionFulu
		v.Fulu = req
		return nil

	case spec.DataVersionElectra:
		req := new(bidadjustment.ElectraAdjustableSubmitBlockRequest)
		if err := req.UnmarshalSSZ(data); err != nil {
			return err
		}
		v.Version = spec.DataVersionElectra
		v.Electra = req
		return nil

	case spec.DataVersionDeneb:
		req := new(bidadjustment.DenebAdjustableSubmitBlockRequest)
		if err := req.UnmarshalSSZ(data); err != nil {
			return err
		}
		v.Version = spec.DataVersionDeneb
		v.Deneb = req
		return nil

	default:
		return errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%s is not supported", fork))
	}
}

func (v *VersionedAdjustableSubmitBlockRequest) UnmarshalSSZ(data []byte) error {
	var err error

	fuluAdjustableSubmitBlockRequest := new(bidadjustment.FuluAdjustableSubmitBlockRequest)
	if err = fuluAdjustableSubmitBlockRequest.UnmarshalSSZ(data); err == nil {
		v.Version = spec.DataVersionFulu
		v.Fulu = fuluAdjustableSubmitBlockRequest
		return nil
	}

	electraAdjustableSubmitBlockRequest := new(bidadjustment.ElectraAdjustableSubmitBlockRequest)
	if err = electraAdjustableSubmitBlockRequest.UnmarshalSSZ(data); err == nil {
		v.Version = spec.DataVersionElectra
		v.Electra = electraAdjustableSubmitBlockRequest
		return nil
	}

	denebAdjustableSubmitBlockRequest := new(bidadjustment.DenebAdjustableSubmitBlockRequest)
	if err = denebAdjustableSubmitBlockRequest.UnmarshalSSZ(data); err == nil {
		v.Version = spec.DataVersionDeneb
		v.Deneb = denebAdjustableSubmitBlockRequest
		return nil
	}

	return errors.Wrap(err, "failed to unmarshal AdjustableSubmitBlockRequest SSZ")
}

func (v *VersionedAdjustableSubmitBlockRequest) MarshalSSZ() ([]byte, error) {
	switch v.Version {
	case spec.DataVersionFulu:
		return v.Fulu.MarshalSSZ()
	case spec.DataVersionElectra:
		return v.Electra.MarshalSSZ()
	case spec.DataVersionDeneb:
		return v.Deneb.MarshalSSZ()
	}
	return nil, errors.New("unknown data version")
}

func RelayGrpcHeaderSubmissionToVersioned(header *relaygrpc.StreamHeaderResponse, URL []byte, forkVersion spec.DataVersion) (*HeaderSubmissionV3, error) {
	if header == nil {
		return nil, errors.New("nil struct")
	}
	if header.BidTrace == nil || header.ExecutionPayloadHeader == nil {
		return nil, errors.New("no bid trace or execution payload header")
	}
	switch forkVersion {
	case spec.DataVersionFulu:
		fuluSubmission, err := relaygrpc.ProtoRequestToFuluHeaderSubmission(header)
		if err != nil {
			return nil, err
		}
		return RelaygrpcFuluHeaderSubmissionToVersioned(fuluSubmission, URL, header.GetTxCount()), nil
	case spec.DataVersionElectra:
		electraSubmission, err := relaygrpc.ProtoRequestToElectraHeaderSubmission(header)
		if err != nil {
			return nil, err
		}
		return RelaygrpcElectraHeaderSubmissionToVersioned(electraSubmission, URL, header.GetTxCount()), nil
	default:
	}
	denebSubmission, err := relaygrpc.ProtoRequestToDenebHeaderSubmission(header)
	if err != nil {
		return nil, err
	}
	return RelaygrpcDenebHeaderSubmissionToVersioned(denebSubmission, URL, header.GetTxCount()), nil
}

func RelaygrpcDenebHeaderSubmissionToVersioned(grpcSubmission *relaygrpc.SignedHeaderSubmissionDeneb, URL []byte, txCount uint64) *HeaderSubmissionV3 {
	submission := &VersionedSignedHeaderSubmission{
		Version: spec.DataVersionDeneb,
		Deneb: &SignedHeaderSubmissionDeneb{
			Message: HeaderSubmissionDenebV2{
				BidTrace:               grpcSubmission.Message.BidTrace,
				ExecutionPayloadHeader: grpcSubmission.Message.ExecutionPayloadHeader,
				Commitments:            grpcSubmission.Message.Commitments,
			},
			Signature: grpcSubmission.Signature,
		},
	}
	return &HeaderSubmissionV3{
		URL:        URL,
		TxCount:    uint32(txCount),
		Submission: submission,
	}
}
func RelaygrpcElectraHeaderSubmissionToVersioned(grpcSubmission *relaygrpc.SignedHeaderSubmissionElectra, URL []byte, txCount uint64) *HeaderSubmissionV3 {
	submission := &VersionedSignedHeaderSubmission{
		Version: spec.DataVersionElectra,
		Electra: &SignedHeaderSubmissionElectra{
			Message: HeaderSubmissionElectra{
				BidTrace:               grpcSubmission.Message.BidTrace,
				ExecutionPayloadHeader: grpcSubmission.Message.ExecutionPayloadHeader,
				Commitments:            grpcSubmission.Message.Commitments,
				ExecutionRequests:      grpcSubmission.Message.ExecutionRequests,
			},
			Signature: grpcSubmission.Signature,
		},
	}
	return &HeaderSubmissionV3{
		URL:        URL,
		TxCount:    uint32(txCount),
		Submission: submission,
	}
}
func RelaygrpcFuluHeaderSubmissionToVersioned(grpcSubmission *relaygrpc.SignedHeaderSubmissionFulu, URL []byte, txCount uint64) *HeaderSubmissionV3 {
	submission := &VersionedSignedHeaderSubmission{
		Version: spec.DataVersionFulu,
		Fulu: &SignedHeaderSubmissionFulu{
			Message: HeaderSubmissionFulu{
				BidTrace:               grpcSubmission.Message.BidTrace,
				ExecutionPayloadHeader: grpcSubmission.Message.ExecutionPayloadHeader,
				Commitments:            grpcSubmission.Message.Commitments,
				ExecutionRequests:      grpcSubmission.Message.ExecutionRequests,
			},
			Signature: grpcSubmission.Signature,
		},
	}
	return &HeaderSubmissionV3{
		URL:        URL,
		TxCount:    uint32(txCount),
		Submission: submission,
	}
}

func BuilderBlockRequestToSignedBuilderBidV3(payload *HeaderSubmissionV3, sk *bls.SecretKey, pubkey *phase0.BLSPubKey, domain phase0.Domain) (*builderSpec.VersionedSignedBuilderBid, error) {

	switch payload.Submission.Version { //nolint:exhaustive
	case spec.DataVersionDeneb:
		in := payload.Submission.Deneb.Message.Commitments
		KZGCommitments := make([]deneb.KZGCommitment, len(in))
		for i, c := range in {
			KZGCommitments[i] = deneb.KZGCommitment(c)
		}
		builderBid := builderApiDeneb.BuilderBid{
			Header:             payload.Submission.Deneb.Message.ExecutionPayloadHeader,
			BlobKZGCommitments: KZGCommitments,
			Value:              payload.Submission.Deneb.Message.BidTrace.Value,
			Pubkey:             *pubkey,
		}

		sig, err := ssz.SignMessage(&builderBid, domain, sk)
		if err != nil {
			return nil, err
		}

		return &builderSpec.VersionedSignedBuilderBid{
			Version: spec.DataVersionDeneb,
			Deneb: &builderApiDeneb.SignedBuilderBid{
				Message:   &builderBid,
				Signature: sig,
			},
		}, nil
	case spec.DataVersionElectra:
		in := payload.Submission.Electra.Message.Commitments
		KZGCommitments := make([]deneb.KZGCommitment, len(in))
		for i, c := range in {
			KZGCommitments[i] = deneb.KZGCommitment(c)
		}
		builderBid := builderApiElectra.BuilderBid{
			Header:             payload.Submission.Electra.Message.ExecutionPayloadHeader,
			BlobKZGCommitments: KZGCommitments,
			Value:              payload.Submission.Electra.Message.BidTrace.Value,
			Pubkey:             *pubkey,
			ExecutionRequests:  payload.Submission.Electra.Message.ExecutionRequests,
		}
		sig, err := ssz.SignMessage(&builderBid, domain, sk)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSignedBuilderBid{
			Version: spec.DataVersionElectra,
			Electra: &builderApiElectra.SignedBuilderBid{
				Message:   &builderBid,
				Signature: sig,
			},
		}, nil

	case spec.DataVersionFulu:
		// The BuilderBid type for fulu is the same as that of electra
		in := payload.Submission.Fulu.Message.Commitments
		KZGCommitments := make([]deneb.KZGCommitment, len(in))
		for i, c := range in {
			KZGCommitments[i] = deneb.KZGCommitment(c)
		}
		builderBid := builderApiElectra.BuilderBid{
			Header:             payload.Submission.Fulu.Message.ExecutionPayloadHeader,
			BlobKZGCommitments: KZGCommitments,
			Value:              payload.Submission.Fulu.Message.BidTrace.Value,
			Pubkey:             *pubkey,
			ExecutionRequests:  payload.Submission.Fulu.Message.ExecutionRequests,
		}
		sig, err := ssz.SignMessage(&builderBid, domain, sk)
		if err != nil {
			return nil, err
		}
		return &builderSpec.VersionedSignedBuilderBid{
			Version: spec.DataVersionFulu,
			Fulu: &builderApiElectra.SignedBuilderBid{
				Message:   &builderBid,
				Signature: sig,
			},
		}, nil

	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%s is not supported", payload.Submission.Version))
	}
}
