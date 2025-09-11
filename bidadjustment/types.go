package bidadjustment

// sourced from https://github.com/blombern/builder/tree/deneb-adjusting
import (
	d "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	electra "github.com/attestantio/go-eth2-client/spec/electra"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
)

type AdjustmentData struct {
	StateRoot               [32]byte `ssz-size:"32"`
	TransactionsRoot        [32]byte `ssz-size:"32"`
	ReceiptsRoot            [32]byte `ssz-size:"32"`
	BuilderAddress          [20]byte `ssz-size:"20"`
	BuilderProof            [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeeRecipientAddress     [20]byte `ssz-size:"20"`
	FeeRecipientProof       [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeePayerAddress         [20]byte `ssz-size:"20"`
	FeePayerProof           [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	PlaceholderTxProof      [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
	PlaceholderReceiptProof [][]byte `ssz-size:"?,?" ssz-max:"64,1073741824"`
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
	CLPlaceholderTxProof    [][32]byte `ssz-size:"?,32" ssz-max:"64,1073741824"`
	PlaceholderReceiptProof [][]byte   `ssz-size:"?,?" ssz-max:"64,1073741824"`
	PrePaymentLogsBloom     [256]byte  `ssz-size:"256"`
}

type DenebAdjustableSubmitBlockRequest struct {
	Message          *v1.BidTrace
	ExecutionPayload *deneb.ExecutionPayload
	BlobsBundle      *d.BlobsBundle
	Signature        phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData   *AdjustmentData
}

type ElectraAdjustableSubmitBlockRequest struct {
	Message           *v1.BidTrace
	ExecutionPayload  *deneb.ExecutionPayload
	BlobsBundle       *d.BlobsBundle
	ExecutionRequests *electra.ExecutionRequests
	Signature         phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData    *AdjustmentData
}
