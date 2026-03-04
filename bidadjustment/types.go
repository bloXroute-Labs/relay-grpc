package bidadjustment

// sourced from https://github.com/blombern/builder/tree/deneb-adjusting
import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	d "github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-builder-client/api/fulu"
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
	AdjustmentDataVersion3
)

// VersionedAdjustmentData versioned adjustment data.
type VersionedAdjustmentData struct {
	Version AdjustmentDataVersion
	V1      *AdjustmentData
	V2      *AdjustmentDataV2
	V3      *AdjustmentDataV3
}

// IsEmpty returns true if there is no adjustment data.
func (v *VersionedAdjustmentData) IsEmpty() bool {
	switch v.Version {
	case AdjustmentDataVersion3:
		return v.V3 == nil
	case AdjustmentDataVersion2:
		return v.V2 == nil
	case AdjustmentDataVersion1:
		return v.V1 == nil
	default:
		return true
	}
}

func (v *VersionedAdjustmentData) MarshalJSON() ([]byte, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion3:
		return json.Marshal(v.V3)
	case AdjustmentDataVersion2:
		return json.Marshal(v.V2)
	case AdjustmentDataVersion1:
		return json.Marshal(v.V1)
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func jsonUnmarshalDisallowUnknownFields(input []byte, v any) error {
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}

func (v *VersionedAdjustmentData) UnmarshalJSON(input []byte) error {
	dataV3 := new(AdjustmentDataV3)
	if err := jsonUnmarshalDisallowUnknownFields(input, dataV3); err == nil {
		v.Version = AdjustmentDataVersion3
		v.V3 = dataV3
		return nil
	}

	dataV2 := new(AdjustmentDataV2)
	if err := jsonUnmarshalDisallowUnknownFields(input, dataV2); err == nil {
		v.Version = AdjustmentDataVersion2
		v.V2 = dataV2
		return nil
	}

	dataV1 := new(AdjustmentData)
	if err := jsonUnmarshalDisallowUnknownFields(input, dataV1); err == nil {
		v.Version = AdjustmentDataVersion1
		v.V1 = dataV1
		return nil
	}

	return errors.Wrapf(ErrInvalidVersion, "failed to unmarshal VersionedAdjustmentData JSON")
}

func (v *VersionedAdjustmentData) BuilderAddress() (ethCommon.Address, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion3:
		return v.V3.BuilderAddress, nil
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
	case AdjustmentDataVersion3:
		return v.V3.FeeRecipientAddress, nil
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
	case AdjustmentDataVersion3:
		return v.V3.FeePayerAddress, nil
	case AdjustmentDataVersion2:
		return v.V2.FeePayerAddress, nil
	case AdjustmentDataVersion1:
		return v.V1.FeePayerAddress, nil
	default:
		return ethCommon.Address{}, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) PlaceholderTransactionProof() ([][]byte, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion3:
		return v.V3.ELPlaceholderTransactionProof, nil
	case AdjustmentDataVersion2:
		return v.V2.ELPlaceholderTransactionProof, nil
	case AdjustmentDataVersion1:
		return v.V1.PlaceholderTransactionProof, nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

func (v *VersionedAdjustmentData) PlaceholderReceiptProof() ([][]byte, error) {
	switch v.Version { //nolint:exhaustive
	case AdjustmentDataVersion3:
		return v.V3.ELPlaceholderReceiptProof, nil
	case AdjustmentDataVersion2:
		return v.V2.PlaceholderReceiptProof, nil
	case AdjustmentDataVersion1:
		return v.V1.PlaceholderReceiptProof, nil
	default:
		return nil, errors.Wrap(ErrInvalidVersion, fmt.Sprintf("%d is not supported", v.Version))
	}
}

type AdjustmentData struct {
	StateRoot                   [32]byte `json:"state_root" ssz-size:"32"`
	TransactionsRoot            [32]byte `json:"transactions_root" ssz-size:"32"`
	ReceiptsRoot                [32]byte `json:"receipts_root" ssz-size:"32"`
	BuilderAddress              [20]byte `json:"builder_address" ssz-size:"20"`
	BuilderProof                [][]byte `json:"builder_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeeRecipientAddress         [20]byte `json:"fee_recipient_address" ssz-size:"20"`
	FeeRecipientProof           [][]byte `json:"fee_recipient_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeePayerAddress             [20]byte `json:"fee_payer_address" ssz-size:"20"`
	FeePayerProof               [][]byte `json:"fee_payer_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	PlaceholderTransactionProof [][]byte `json:"placeholder_transaction_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	PlaceholderReceiptProof     [][]byte `json:"placeholder_receipt_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
}

// MarshalJSON implements json.Marshaler for AdjustmentData
func (a *AdjustmentData) MarshalJSON() ([]byte, error) {
	type Alias struct {
		StateRoot                   string   `json:"state_root"`
		TransactionsRoot            string   `json:"transactions_root"`
		ReceiptsRoot                string   `json:"receipts_root"`
		BuilderAddress              string   `json:"builder_address"`
		BuilderProof                []string `json:"builder_proof"`
		FeeRecipientAddress         string   `json:"fee_recipient_address"`
		FeeRecipientProof           []string `json:"fee_recipient_proof"`
		FeePayerAddress             string   `json:"fee_payer_address"`
		FeePayerProof               []string `json:"fee_payer_proof"`
		PlaceholderTransactionProof []string `json:"placeholder_transaction_proof"`
		PlaceholderReceiptProof     []string `json:"placeholder_receipt_proof"`
	}

	return json.Marshal(&Alias{
		StateRoot:                   "0x" + hex.EncodeToString(a.StateRoot[:]),
		TransactionsRoot:            "0x" + hex.EncodeToString(a.TransactionsRoot[:]),
		ReceiptsRoot:                "0x" + hex.EncodeToString(a.ReceiptsRoot[:]),
		BuilderAddress:              "0x" + hex.EncodeToString(a.BuilderAddress[:]),
		BuilderProof:                bytesArrayToHexStrings(a.BuilderProof),
		FeeRecipientAddress:         "0x" + hex.EncodeToString(a.FeeRecipientAddress[:]),
		FeeRecipientProof:           bytesArrayToHexStrings(a.FeeRecipientProof),
		FeePayerAddress:             "0x" + hex.EncodeToString(a.FeePayerAddress[:]),
		FeePayerProof:               bytesArrayToHexStrings(a.FeePayerProof),
		PlaceholderTransactionProof: bytesArrayToHexStrings(a.PlaceholderTransactionProof),
		PlaceholderReceiptProof:     bytesArrayToHexStrings(a.PlaceholderReceiptProof),
	})
}

// UnmarshalJSON implements json.Unmarshaler for AdjustmentData
func (a *AdjustmentData) UnmarshalJSON(data []byte) error {
	type Alias struct {
		StateRoot                   any `json:"state_root"`
		TransactionsRoot            any `json:"transactions_root"`
		ReceiptsRoot                any `json:"receipts_root"`
		BuilderAddress              any `json:"builder_address"`
		BuilderProof                any `json:"builder_proof"`
		FeeRecipientAddress         any `json:"fee_recipient_address"`
		FeeRecipientProof           any `json:"fee_recipient_proof"`
		FeePayerAddress             any `json:"fee_payer_address"`
		FeePayerProof               any `json:"fee_payer_proof"`
		PlaceholderTransactionProof any `json:"placeholder_transaction_proof"`
		PlaceholderReceiptProof     any `json:"placeholder_receipt_proof"`
	}

	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error

	// Parse fixed-size byte arrays (support both hex/base64 strings and number arrays)
	if a.StateRoot, err = parseBytes32(aux.StateRoot); err != nil {
		return fmt.Errorf("state_root: %w", err)
	}
	if a.TransactionsRoot, err = parseBytes32(aux.TransactionsRoot); err != nil {
		return fmt.Errorf("transactions_root: %w", err)
	}
	if a.ReceiptsRoot, err = parseBytes32(aux.ReceiptsRoot); err != nil {
		return fmt.Errorf("receipts_root: %w", err)
	}
	if a.BuilderAddress, err = parseBytes20(aux.BuilderAddress); err != nil {
		return fmt.Errorf("builder_address: %w", err)
	}
	if a.FeeRecipientAddress, err = parseBytes20(aux.FeeRecipientAddress); err != nil {
		return fmt.Errorf("fee_recipient_address: %w", err)
	}
	if a.FeePayerAddress, err = parseBytes20(aux.FeePayerAddress); err != nil {
		return fmt.Errorf("fee_payer_address: %w", err)
	}

	// Parse dynamic byte arrays (support both hex/base64 string arrays and number arrays)
	if a.BuilderProof, err = parseBytesArray(aux.BuilderProof); err != nil {
		return fmt.Errorf("builder_proof: %w", err)
	}
	if a.FeeRecipientProof, err = parseBytesArray(aux.FeeRecipientProof); err != nil {
		return fmt.Errorf("fee_recipient_proof: %w", err)
	}
	if a.FeePayerProof, err = parseBytesArray(aux.FeePayerProof); err != nil {
		return fmt.Errorf("fee_payer_proof: %w", err)
	}
	if a.PlaceholderTransactionProof, err = parseBytesArray(aux.PlaceholderTransactionProof); err != nil {
		return fmt.Errorf("placeholder_transaction_proof: %w", err)
	}
	if a.PlaceholderReceiptProof, err = parseBytesArray(aux.PlaceholderReceiptProof); err != nil {
		return fmt.Errorf("placeholder_receipt_proof: %w", err)
	}

	return nil
}

type AdjustmentDataV2 struct {
	ELTransactionsRoot            [32]byte   `json:"el_transactions_root" ssz-size:"32"`
	ELWithdrawalsRoot             [32]byte   `json:"el_withdrawals_root" ssz-size:"32"`
	BuilderAddress                [20]byte   `json:"builder_address" ssz-size:"20"`
	BuilderProof                  [][]byte   `json:"builder_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeeRecipientAddress           [20]byte   `json:"fee_recipient_address" ssz-size:"20"`
	FeeRecipientProof             [][]byte   `json:"fee_recipient_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeePayerAddress               [20]byte   `json:"fee_payer_address" ssz-size:"20"`
	FeePayerProof                 [][]byte   `json:"fee_payer_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	ELPlaceholderTransactionProof [][]byte   `json:"el_placeholder_transaction_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	CLPlaceholderTransactionProof [][32]byte `json:"cl_placeholder_transaction_proof" ssz-size:"?,32" ssz-max:"64,1073741824"`
	PlaceholderReceiptProof       [][]byte   `json:"placeholder_receipt_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	PrePaymentLogsBloom           [256]byte  `json:"pre_payment_logs_bloom" ssz-size:"256"`
}

// MarshalJSON implements json.Marshaler for AdjustmentDataV2
func (a *AdjustmentDataV2) MarshalJSON() ([]byte, error) {
	type Alias struct {
		ELTransactionsRoot            string   `json:"el_transactions_root"`
		ELWithdrawalsRoot             string   `json:"el_withdrawals_root"`
		BuilderAddress                string   `json:"builder_address"`
		BuilderProof                  []string `json:"builder_proof"`
		FeeRecipientAddress           string   `json:"fee_recipient_address"`
		FeeRecipientProof             []string `json:"fee_recipient_proof"`
		FeePayerAddress               string   `json:"fee_payer_address"`
		FeePayerProof                 []string `json:"fee_payer_proof"`
		ELPlaceholderTransactionProof []string `json:"el_placeholder_transaction_proof"`
		CLPlaceholderTransactionProof []string `json:"cl_placeholder_transaction_proof"`
		PlaceholderReceiptProof       []string `json:"placeholder_receipt_proof"`
		PrePaymentLogsBloom           string   `json:"pre_payment_logs_bloom"`
	}

	return json.Marshal(&Alias{
		ELTransactionsRoot:            "0x" + hex.EncodeToString(a.ELTransactionsRoot[:]),
		ELWithdrawalsRoot:             "0x" + hex.EncodeToString(a.ELWithdrawalsRoot[:]),
		BuilderAddress:                "0x" + hex.EncodeToString(a.BuilderAddress[:]),
		BuilderProof:                  bytesArrayToHexStrings(a.BuilderProof),
		FeeRecipientAddress:           "0x" + hex.EncodeToString(a.FeeRecipientAddress[:]),
		FeeRecipientProof:             bytesArrayToHexStrings(a.FeeRecipientProof),
		FeePayerAddress:               "0x" + hex.EncodeToString(a.FeePayerAddress[:]),
		FeePayerProof:                 bytesArrayToHexStrings(a.FeePayerProof),
		ELPlaceholderTransactionProof: bytesArrayToHexStrings(a.ELPlaceholderTransactionProof),
		CLPlaceholderTransactionProof: bytes32ArrayToHexStrings(a.CLPlaceholderTransactionProof),
		PlaceholderReceiptProof:       bytesArrayToHexStrings(a.PlaceholderReceiptProof),
		PrePaymentLogsBloom:           "0x" + hex.EncodeToString(a.PrePaymentLogsBloom[:]),
	})
}

// UnmarshalJSON implements json.Unmarshaler for AdjustmentDataV2
func (a *AdjustmentDataV2) UnmarshalJSON(data []byte) error {
	type Alias struct {
		ELTransactionsRoot            any `json:"el_transactions_root"`
		ELWithdrawalsRoot             any `json:"el_withdrawals_root"`
		BuilderAddress                any `json:"builder_address"`
		BuilderProof                  any `json:"builder_proof"`
		FeeRecipientAddress           any `json:"fee_recipient_address"`
		FeeRecipientProof             any `json:"fee_recipient_proof"`
		FeePayerAddress               any `json:"fee_payer_address"`
		FeePayerProof                 any `json:"fee_payer_proof"`
		ELPlaceholderTransactionProof any `json:"el_placeholder_transaction_proof"`
		CLPlaceholderTransactionProof any `json:"cl_placeholder_transaction_proof"`
		PlaceholderReceiptProof       any `json:"placeholder_receipt_proof"`
		PrePaymentLogsBloom           any `json:"pre_payment_logs_bloom"`
	}

	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error

	// Parse fixed-size byte arrays
	if a.ELTransactionsRoot, err = parseBytes32(aux.ELTransactionsRoot); err != nil {
		return fmt.Errorf("el_transactions_root: %w", err)
	}
	if a.ELWithdrawalsRoot, err = parseBytes32(aux.ELWithdrawalsRoot); err != nil {
		return fmt.Errorf("el_withdrawals_root: %w", err)
	}
	if a.BuilderAddress, err = parseBytes20(aux.BuilderAddress); err != nil {
		return fmt.Errorf("builder_address: %w", err)
	}
	if a.FeeRecipientAddress, err = parseBytes20(aux.FeeRecipientAddress); err != nil {
		return fmt.Errorf("fee_recipient_address: %w", err)
	}
	if a.FeePayerAddress, err = parseBytes20(aux.FeePayerAddress); err != nil {
		return fmt.Errorf("fee_payer_address: %w", err)
	}
	if a.PrePaymentLogsBloom, err = parseBytes256(aux.PrePaymentLogsBloom); err != nil {
		return fmt.Errorf("pre_payment_logs_bloom: %w", err)
	}

	// Parse dynamic byte arrays
	if a.BuilderProof, err = parseBytesArray(aux.BuilderProof); err != nil {
		return fmt.Errorf("builder_proof: %w", err)
	}
	if a.FeeRecipientProof, err = parseBytesArray(aux.FeeRecipientProof); err != nil {
		return fmt.Errorf("fee_recipient_proof: %w", err)
	}
	if a.FeePayerProof, err = parseBytesArray(aux.FeePayerProof); err != nil {
		return fmt.Errorf("fee_payer_proof: %w", err)
	}
	if a.ELPlaceholderTransactionProof, err = parseBytesArray(aux.ELPlaceholderTransactionProof); err != nil {
		return fmt.Errorf("el_placeholder_transaction_proof: %w", err)
	}
	if a.PlaceholderReceiptProof, err = parseBytesArray(aux.PlaceholderReceiptProof); err != nil {
		return fmt.Errorf("placeholder_receipt_proof: %w", err)
	}
	if a.CLPlaceholderTransactionProof, err = parseBytes32Array(aux.CLPlaceholderTransactionProof); err != nil {
		return fmt.Errorf("cl_placeholder_transaction_proof: %w", err)
	}

	return nil
}

type AdjustmentDataV3 struct {
	ELTransactionsRoot            [32]byte   `json:"el_transactions_root" ssz-size:"32"`
	ELWithdrawalsRoot             [32]byte   `json:"el_withdrawals_root" ssz-size:"32"`
	BuilderAddress                [20]byte   `json:"builder_address" ssz-size:"20"`
	BuilderProof                  [][]byte   `json:"builder_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeeRecipientAddress           [20]byte   `json:"fee_recipient_address" ssz-size:"20"`
	FeeRecipientProof             [][]byte   `json:"fee_recipient_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	FeePayerAddress               [20]byte   `json:"fee_payer_address" ssz-size:"20"`
	FeePayerProof                 [][]byte   `json:"fee_payer_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	ELPlaceholderTransactionProof [][]byte   `json:"el_placeholder_transaction_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	CLPlaceholderTransactionProof [][32]byte `json:"cl_placeholder_transaction_proof" ssz-size:"?,32" ssz-max:"64,1073741824"`
	ELPlaceholderReceiptProof     [][]byte   `json:"el_placeholder_receipt_proof" ssz-size:"?,?" ssz-max:"64,1073741824"`
	PrePaymentLogsBloom           [256]byte  `json:"pre_payment_logs_bloom" ssz-size:"256"`
	PlaceholderGasUsed            uint64     `json:"placeholder_gas_used"`
}

// MarshalJSON implements json.Marshaler for AdjustmentDataV3
func (a *AdjustmentDataV3) MarshalJSON() ([]byte, error) {
	type Alias struct {
		ELTransactionsRoot            string   `json:"el_transactions_root"`
		ELWithdrawalsRoot             string   `json:"el_withdrawals_root"`
		BuilderAddress                string   `json:"builder_address"`
		BuilderProof                  []string `json:"builder_proof"`
		FeeRecipientAddress           string   `json:"fee_recipient_address"`
		FeeRecipientProof             []string `json:"fee_recipient_proof"`
		FeePayerAddress               string   `json:"fee_payer_address"`
		FeePayerProof                 []string `json:"fee_payer_proof"`
		ELPlaceholderTransactionProof []string `json:"el_placeholder_transaction_proof"`
		CLPlaceholderTransactionProof []string `json:"cl_placeholder_transaction_proof"`
		ELPlaceholderReceiptProof     []string `json:"el_placeholder_receipt_proof"`
		PrePaymentLogsBloom           string   `json:"pre_payment_logs_bloom"`
		PlaceholderGasUsed            uint64   `json:"placeholder_gas_used"`
	}

	return json.Marshal(&Alias{
		ELTransactionsRoot:            "0x" + hex.EncodeToString(a.ELTransactionsRoot[:]),
		ELWithdrawalsRoot:             "0x" + hex.EncodeToString(a.ELWithdrawalsRoot[:]),
		BuilderAddress:                "0x" + hex.EncodeToString(a.BuilderAddress[:]),
		BuilderProof:                  bytesArrayToHexStrings(a.BuilderProof),
		FeeRecipientAddress:           "0x" + hex.EncodeToString(a.FeeRecipientAddress[:]),
		FeeRecipientProof:             bytesArrayToHexStrings(a.FeeRecipientProof),
		FeePayerAddress:               "0x" + hex.EncodeToString(a.FeePayerAddress[:]),
		FeePayerProof:                 bytesArrayToHexStrings(a.FeePayerProof),
		ELPlaceholderTransactionProof: bytesArrayToHexStrings(a.ELPlaceholderTransactionProof),
		CLPlaceholderTransactionProof: bytes32ArrayToHexStrings(a.CLPlaceholderTransactionProof),
		ELPlaceholderReceiptProof:     bytesArrayToHexStrings(a.ELPlaceholderReceiptProof),
		PrePaymentLogsBloom:           "0x" + hex.EncodeToString(a.PrePaymentLogsBloom[:]),
		PlaceholderGasUsed:            a.PlaceholderGasUsed,
	})
}

// UnmarshalJSON implements json.Unmarshaler for AdjustmentDataV3
func (a *AdjustmentDataV3) UnmarshalJSON(data []byte) error {
	type Alias struct {
		ELTransactionsRoot            any    `json:"el_transactions_root"`
		ELWithdrawalsRoot             any    `json:"el_withdrawals_root"`
		BuilderAddress                any    `json:"builder_address"`
		BuilderProof                  any    `json:"builder_proof"`
		FeeRecipientAddress           any    `json:"fee_recipient_address"`
		FeeRecipientProof             any    `json:"fee_recipient_proof"`
		FeePayerAddress               any    `json:"fee_payer_address"`
		FeePayerProof                 any    `json:"fee_payer_proof"`
		ELPlaceholderTransactionProof any    `json:"el_placeholder_transaction_proof"`
		CLPlaceholderTransactionProof any    `json:"cl_placeholder_transaction_proof"`
		ELPlaceholderReceiptProof     any    `json:"el_placeholder_receipt_proof"`
		PrePaymentLogsBloom           any    `json:"pre_payment_logs_bloom"`
		PlaceholderGasUsed            uint64 `json:"placeholder_gas_used"`
	}

	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error

	// Parse fixed-size byte arrays
	if a.ELTransactionsRoot, err = parseBytes32(aux.ELTransactionsRoot); err != nil {
		return fmt.Errorf("el_transactions_root: %w", err)
	}
	if a.ELWithdrawalsRoot, err = parseBytes32(aux.ELWithdrawalsRoot); err != nil {
		return fmt.Errorf("el_withdrawals_root: %w", err)
	}
	if a.BuilderAddress, err = parseBytes20(aux.BuilderAddress); err != nil {
		return fmt.Errorf("builder_address: %w", err)
	}
	if a.FeeRecipientAddress, err = parseBytes20(aux.FeeRecipientAddress); err != nil {
		return fmt.Errorf("fee_recipient_address: %w", err)
	}
	if a.FeePayerAddress, err = parseBytes20(aux.FeePayerAddress); err != nil {
		return fmt.Errorf("fee_payer_address: %w", err)
	}
	if a.PrePaymentLogsBloom, err = parseBytes256(aux.PrePaymentLogsBloom); err != nil {
		return fmt.Errorf("pre_payment_logs_bloom: %w", err)
	}

	// Parse dynamic byte arrays
	if a.BuilderProof, err = parseBytesArray(aux.BuilderProof); err != nil {
		return fmt.Errorf("builder_proof: %w", err)
	}
	if a.FeeRecipientProof, err = parseBytesArray(aux.FeeRecipientProof); err != nil {
		return fmt.Errorf("fee_recipient_proof: %w", err)
	}
	if a.FeePayerProof, err = parseBytesArray(aux.FeePayerProof); err != nil {
		return fmt.Errorf("fee_payer_proof: %w", err)
	}
	if a.ELPlaceholderTransactionProof, err = parseBytesArray(aux.ELPlaceholderTransactionProof); err != nil {
		return fmt.Errorf("el_placeholder_transaction_proof: %w", err)
	}
	if a.ELPlaceholderReceiptProof, err = parseBytesArray(aux.ELPlaceholderReceiptProof); err != nil {
		return fmt.Errorf("el_placeholder_receipt_proof: %w", err)
	}
	if a.CLPlaceholderTransactionProof, err = parseBytes32Array(aux.CLPlaceholderTransactionProof); err != nil {
		return fmt.Errorf("cl_placeholder_transaction_proof: %w", err)
	}

	a.PlaceholderGasUsed = aux.PlaceholderGasUsed

	return nil
}

type DenebAdjustableSubmitBlockRequest struct {
	Message          *v1.BidTrace
	ExecutionPayload *deneb.ExecutionPayload
	BlobsBundle      *d.BlobsBundle
	Signature        phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData   *VersionedAdjustmentData
}

type ElectraAdjustableSubmitBlockRequest struct {
	Message           *v1.BidTrace
	ExecutionPayload  *deneb.ExecutionPayload
	BlobsBundle       *d.BlobsBundle
	ExecutionRequests *electra.ExecutionRequests
	Signature         phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData    *VersionedAdjustmentData
}

type FuluAdjustableSubmitBlockRequest struct {
	Message           *v1.BidTrace
	ExecutionPayload  *deneb.ExecutionPayload
	BlobsBundle       *fulu.BlobsBundle
	ExecutionRequests *electra.ExecutionRequests
	Signature         phase0.BLSSignature `ssz-size:"96"`
	AdjustmentData    *VersionedAdjustmentData
}

// TODO: hopefully won't need versioning functions after this PR is done
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

func AdjustmentDataV3ToVersioned(adjustmentData *AdjustmentDataV3) *VersionedAdjustmentData {
	if adjustmentData == nil {
		return nil
	}

	return &VersionedAdjustmentData{
		Version: AdjustmentDataVersion3,
		V3:      adjustmentData,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// bytesArrayToHexStrings converts [][]byte to []string with hex encoding
func bytesArrayToHexStrings(data [][]byte) []string {
	if data == nil {
		return nil
	}
	result := make([]string, len(data))
	for i, b := range data {
		result[i] = "0x" + hex.EncodeToString(b)
	}
	return result
}

// parseBytes32 parses interface{} to [32]byte, supporting both hex/base64 strings and number arrays
func parseBytes32(v any) ([32]byte, error) {
	var result [32]byte

	switch val := v.(type) {
	case string:
		decoded, err := decodeHexOrBase64(val, 32)
		if err != nil {
			return result, err
		}
		copy(result[:], decoded)
	case []any:
		if len(val) != 32 {
			return result, fmt.Errorf("expected 32 bytes, got %d", len(val))
		}
		for i, num := range val {
			switch n := num.(type) {
			case float64:
				result[i] = byte(n)
			default:
				return result, fmt.Errorf("invalid byte at index %d", i)
			}
		}
	default:
		return result, fmt.Errorf("expected string or array, got %T", v)
	}

	return result, nil
}

// parseBytes20 parses interface{} to [20]byte, supporting both hex/base64 strings and number arrays
func parseBytes20(v any) ([20]byte, error) {
	var result [20]byte

	switch val := v.(type) {
	case string:
		decoded, err := decodeHexOrBase64(val, 20)
		if err != nil {
			return result, err
		}
		copy(result[:], decoded)
	case []any:
		if len(val) != 20 {
			return result, fmt.Errorf("expected 20 bytes, got %d", len(val))
		}
		for i, num := range val {
			switch n := num.(type) {
			case float64:
				result[i] = byte(n)
			default:
				return result, fmt.Errorf("invalid byte at index %d", i)
			}
		}
	default:
		return result, fmt.Errorf("expected string or array, got %T", v)
	}

	return result, nil
}

// parseBytesArray parses interface{} to [][]byte, supporting hex/base64 string arrays and number arrays
func parseBytesArray(v any) ([][]byte, error) {
	if v == nil {
		return nil, nil
	}

	switch val := v.(type) {
	case []any:
		result := make([][]byte, len(val))
		for i, item := range val {
			switch itemVal := item.(type) {
			case string:
				decoded, err := decodeHexOrBase64(itemVal, -1)
				if err != nil {
					return nil, fmt.Errorf("item %d: %w", i, err)
				}
				result[i] = decoded
			case []any:
				bytes := make([]byte, len(itemVal))
				for j, num := range itemVal {
					switch n := num.(type) {
					case float64:
						bytes[j] = byte(n)
					default:
						return nil, fmt.Errorf("item %d, byte %d: invalid type %T", i, j, n)
					}
				}
				result[i] = bytes
			default:
				return nil, fmt.Errorf("item %d: expected string or array, got %T", i, itemVal)
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("expected array, got %T", v)
	}
}

// hexDecode decodes a hex string (with or without 0x prefix)
// decodeHexOrBase64 decodes a string as hex if it starts with 0x, otherwise as base64. If length >= 0, enforces length.
func decodeHexOrBase64(s string, length int) ([]byte, error) {
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		decoded, err := hex.DecodeString(s[2:])
		if err != nil {
			return nil, err
		}
		if length >= 0 && len(decoded) != length {
			return nil, fmt.Errorf("expected %d bytes, got %d", length, len(decoded))
		}
		return decoded, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}
	if length >= 0 && len(decoded) != length {
		return nil, fmt.Errorf("expected %d bytes, got %d", length, len(decoded))
	}
	return decoded, nil
}

// parseBytes256 parses interface{} to [256]byte, supporting both hex strings and number arrays
func parseBytes256(v any) ([256]byte, error) {
	var result [256]byte

	switch val := v.(type) {
	case string:
		decoded, err := decodeHexOrBase64(val, 256)
		if err != nil {
			return result, err
		}
		copy(result[:], decoded)
	case []any:
		if len(val) != 256 {
			return result, fmt.Errorf("expected 256 bytes, got %d", len(val))
		}
		for i, num := range val {
			switch n := num.(type) {
			case float64:
				result[i] = byte(n)
			default:
				return result, fmt.Errorf("invalid byte at index %d", i)
			}
		}
	default:
		return result, fmt.Errorf("expected string or array, got %T", v)
	}

	return result, nil
}

// bytes32ArrayToHexStrings converts [][32]byte to []string with hex encoding
func bytes32ArrayToHexStrings(data [][32]byte) []string {
	if data == nil {
		return nil
	}
	result := make([]string, len(data))
	for i, b := range data {
		result[i] = "0x" + hex.EncodeToString(b[:])
	}
	return result
}

// parseBytes32Array parses interface{} to [][32]byte, supporting hex string arrays and number arrays
func parseBytes32Array(v any) ([][32]byte, error) {
	if v == nil {
		return nil, nil
	}

	switch val := v.(type) {
	case []any:
		result := make([][32]byte, len(val))
		for i, item := range val {
			b32, err := parseBytes32(item)
			if err != nil {
				return nil, fmt.Errorf("item %d: %w", i, err)
			}
			result[i] = b32
		}
		return result, nil
	default:
		return nil, fmt.Errorf("expected array, got %T", v)
	}
}
