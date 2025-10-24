package optimisticv3

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
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
	Commitments            [][48]byte                    `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
}

type HeaderSubmissionElectra struct {
	BidTrace               *v1.BidTrace                  `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader `json:"execution_payload_header"`
	ExecutionRequests      *electra.ExecutionRequests    `json:"execution_requests"`
	Commitments            [][48]byte                    `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
	AdjustmentData         AdjustmentDataV2              `json:"adjustment_data"`
}

type HeaderSubmissionFulu struct {
	BidTrace               *v1.BidTrace                  `json:"bid_trace"`
	ExecutionPayloadHeader *deneb.ExecutionPayloadHeader `json:"execution_payload_header"`
	ExecutionRequests      *electra.ExecutionRequests    `json:"execution_requests"`
	Commitments            [][48]byte                    `json:"commitments" ssz-max:"4096" ssz-size:"?,48"`
	AdjustmentData         AdjustmentDataV2              `json:"adjustment_data"`
}

type GetPayloadV3 struct {
	// Hash of the block header from the `SignedHeaderSubmission`.
	BlockHash [32]byte `json:"block_hash" ssz-size:"32"`
	// Timestamp (in milliseconds) when the relay made this request.
	RequestTs uint64 `json:"request_ts"`
	// Bls public key of the signing key that was used to create the `signature` field in `SignedGetPayloadV3`.
	RelayPublicKey [48]byte `json:"relay_public_key" ssz-size:"48"`
}

type SignedGetPayloadV3 struct {
	Message *GetPayloadV3 `json:"message"`
	// Signature from the relay's key that it uses to sign the `get_header` responses.
	Signature phase0.BLSSignature `json:"signature" ssz-size:"96"`
}

type AdjustmentDataV2 struct {
	ELTransactionsRoot      [32]byte   `ssz-size:"32"`
	ELWithdrawalsRoot       [32]byte   `ssz-size:"32"`
	BuilderAddress          [20]byte   `ssz-size:"20"`
	BuilderProof            [][]byte   `ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeeRecipientAddress     [20]byte   `ssz-size:"20"`
	FeeRecipientProof       [][]byte   `ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeePayerAddress         [20]byte   `ssz-size:"20"`
	FeePayerProof           [][]byte   `ssz-size:"?,?" ssz-max:"64,1073741824"`
	ELPlaceholderTxProof    [][]byte   `ssz-size:"?,?" ssz-max:"64,1073741824"`
	CLPlaceholderTxProof    [][32]byte `ssz-size:"?,32" ssz-max:"1073741824"`
	PlaceholderReceiptProof [][]byte   `ssz-size:"?,?" ssz-max:"64,1073741824"`
	PrePaymentLogsBloom     [256]byte  `ssz-size:"256"`
}
