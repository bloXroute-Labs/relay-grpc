package feerecipientanalysis

import (
	"bytes"
	"fmt"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

var TransferGasLimit = uint64(21000)

func SupportedFeeRecipient(isEOA bool, feeRecipientAddress gethcommon.Address, addressByte []byte) bool {
	if isEOA ||
		feeRecipientAddress == FeeRecipient_388C818C ||
		feeRecipientAddress == FeeRecipient_d4E96eF8 ||
		feeRecipientAddress == FeeRecipient_fFEE0878 ||
		feeRecipientAddress == FeeRecipient_eBec795c ||
		bytes.Equal(addressByte, EOA_RocketpoolDistributer) ||
		feeRecipientAddress == FeeRecipient_E978d95B {
		return true
	}
	return false
}

func GetReceiptOutput(isEOA bool, feeRecipientAddress gethcommon.Address, addressByte []byte, adjustedValue *uint256.Int, adjustedValueUint64 uint64, originalTxValue *big.Int, originalBalance *uint256.Int, feePayerAddress gethcommon.Address) (uint64, []*types.Log, []*types.Log, error) {
	var (
		gasUsed             uint64
		adjustedReceiptLog  []*types.Log
		newLogsForLogsBloom []*types.Log
		err                 error
	)

	switch {
	case isEOA:
		gasUsed = TransferGasLimit
		adjustedReceiptLog = []*types.Log{}

	case feeRecipientAddress == FeeRecipient_388C818C:
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom, err = GetReceipt_388C818C(adjustedValue)
		if err != nil {
			return 0, nil, nil, fmt.Errorf("failed to get receipt for fee recipient %s: %v", feeRecipientAddress.Hex(), err)
		}

	case feeRecipientAddress == FeeRecipient_d4E96eF8:
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom = GetReceipt_d4E96eF8()

	case feeRecipientAddress == FeeRecipient_fFEE0878:
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom = GetReceipt_fFEE0878()

	case feeRecipientAddress == FeeRecipient_eBec795c:
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom = GetReceipt_eBec795c()

	case feeRecipientAddress == FeeRecipient_004c2681:
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom, err = GetReceipt_004c2681(feePayerAddress, adjustedValue, []byte{})
		if err != nil {
			return 0, nil, nil, fmt.Errorf("failed to get receipt for fee recipient %s: %v", feeRecipientAddress.Hex(), err)
		}

	case bytes.Equal(addressByte, EOA_RocketpoolDistributer):
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom = GetReceipt_RocketpoolDistributer()

	case feeRecipientAddress == FeeRecipient_E978d95B:
		newBalance := uint256.NewInt(0)
		payment := uint256.NewInt(originalTxValue.Uint64())
		newBalance.Sub(originalBalance, payment)
		adjustedPayment := uint256.NewInt(adjustedValueUint64)
		newBalance.Add(newBalance, adjustedPayment)
		gasUsed, adjustedReceiptLog, newLogsForLogsBloom, err = GetReceipt_E978d95B(newBalance)
		if err != nil {
			return 0, nil, nil, fmt.Errorf("failed to get receipt for fee recipient %s: %v", feeRecipientAddress.Hex(), err)
		}

	default:
		return 0, nil, nil, fmt.Errorf("unsupported fee recipient %s", feeRecipientAddress.Hex())
	}
	return gasUsed, adjustedReceiptLog, newLogsForLogsBloom, nil
}
