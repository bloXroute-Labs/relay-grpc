package optimisticv3

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloXroute-Labs/relay-grpc/bidadjustment"
)

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

type SignedHeaderSubmissionFulu struct {
	Message   HeaderSubmissionFulu `json:"message"`
	Signature phase0.BLSSignature  `json:"signature" ssz-size:"96"`
}

type HeaderSubmissionFulu struct {
	BidTrace               *v1.BidTrace                   `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader  `json:"execution_payload_header"`
	ExecutionRequests      *electra.ExecutionRequests     `json:"execution_requests"`
	Commitments            []deneb.KZGCommitment          `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
	AdjustmentData         bidadjustment.AdjustmentDataV2 `json:"adjustment_data"`
}
