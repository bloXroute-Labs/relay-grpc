package bidadjustment

// sourced from https://github.com/blombern/builder/tree/deneb-adjusting
import (
	"encoding/json"
	"fmt"

	d "github.com/attestantio/go-builder-client/api/deneb"
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var ErrInvalidVersion = errors.New("invalid version")

type AdjustmentDataVersion uint64

const (
	AdjustmentDataVersionUnknown AdjustmentDataVersion = iota
	AdjustmentDataVersion1
	AdjustmentDataVersion2
)

// VersionedAdjustmentData versioned adjustment data.
type VersionedAdjustmentData struct {
	Version AdjustmentDataVersion
	V1      *AdjustmentData
	V2      *AdjustmentDataV2
}

// IsEmpty returns true if there is no adjustment data.
func (v *VersionedAdjustmentData) IsEmpty() bool {
	switch v.Version {
	case AdjustmentDataVersion2:
		return v.V2 == nil
	case AdjustmentDataVersion1:
		return v.V1 == nil
	default:
		return true
	}
}

func (v *VersionedAdjustmentData) MarshalSSZ() ([]byte, error) {
	switch v.Version {
	case AdjustmentDataVersion2:
		return v.V2.MarshalSSZ()
	case AdjustmentDataVersion1:
		return v.V1.MarshalSSZ()
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) UnmarshalSSZ(input []byte) error {
	var err error

	dataV2 := new(AdjustmentDataV2)
	if err = dataV2.UnmarshalSSZ(input); err == nil {
		v.Version = AdjustmentDataVersion2
		v.V2 = dataV2
		return nil
	}

	dataV1 := new(AdjustmentData)
	if err = dataV1.UnmarshalSSZ(input); err == nil {
		v.Version = AdjustmentDataVersion1
		v.V1 = dataV1
		return nil
	}

	return errors.Wrap(err, "failed to unmarshal VersionedAdjustmentData SSZ")
}

func (v *VersionedAdjustmentData) MarshalJSON() ([]byte, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion2:
		return json.Marshal(v.V2)
	case AdjustmentDataVersion1:
		return json.Marshal(v.V1)
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) UnmarshalJSON(input []byte) error {
	var err error

	dataV2 := new(AdjustmentDataV2)
	if err = json.Unmarshal(input, dataV2); err == nil {
		v.Version = AdjustmentDataVersion2
		v.V2 = dataV2
		return nil
	}

	dataV1 := new(AdjustmentData)
	if err = json.Unmarshal(input, dataV1); err == nil {
		v.Version = AdjustmentDataVersion1
		v.V1 = dataV1
		return nil
	}

	return errors.Wrap(err, "failed to unmarshal VersionedAdjustmentData JSON")
}

func (v *VersionedAdjustmentData) BuilderAddress() (ethCommon.Address, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion2:
		return v.V2.BuilderAddress, nil
	case AdjustmentDataVersion1:
		return v.V1.BuilderAddress, nil
	default:
		return ethCommon.Address{}, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) FeeRecipientAddress() (ethCommon.Address, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion2:
		return v.V2.FeeRecipientAddress, nil
	case AdjustmentDataVersion1:
		return v.V1.FeeRecipientAddress, nil
	default:
		return ethCommon.Address{}, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) FeePayerAddress() (ethCommon.Address, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion2:
		return v.V2.FeePayerAddress, nil
	case AdjustmentDataVersion1:
		return v.V1.FeePayerAddress, nil
	default:
		return ethCommon.Address{}, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) PlaceholderTxProof() ([][]byte, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion2:
		return v.V2.ELPlaceholderTxProof, nil
	case AdjustmentDataVersion1:
		return v.V1.PlaceholderTxProof, nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) PlaceholderReceiptProof() ([][]byte, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion2:
		return v.V2.PlaceholderReceiptProof, nil
	case AdjustmentDataVersion1:
		return v.V1.PlaceholderReceiptProof, nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

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

func AdjustmentDataV1ToVersioned(adjustmentData *AdjustmentData) *VersionedAdjustmentData {
	if adjustmentData == nil {
		return nil
	}

	return &VersionedAdjustmentData{
		Version: AdjustmentDataVersion1,
		V1:      adjustmentData,
	}
}

func AdjustmentDataV2ToVersioned(adjustmentData *AdjustmentDataV2) *VersionedAdjustmentData {
	if adjustmentData == nil {
		return nil
	}

	return &VersionedAdjustmentData{
		Version: AdjustmentDataVersion2,
		V2:      adjustmentData,
	}
}
