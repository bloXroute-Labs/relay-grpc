package feerecipientanalysis

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestGetReceipt_388C818C(t *testing.T) {
	value := 22093171402303330
	gas, logs, _, err := GetReceipt_388C818C(uint256.NewInt(uint64(value)))
	require.NoError(t, err)
	if gas != 22111 {
		t.Errorf("Expected gas 22111, got %d", gas)
	}
	if len(logs[0].Topics) != 1 {
		t.Errorf("Expected 1 topic, got %d", len(logs[0].Topics))
	}
	if logs[0].Topics[0] != common.HexToHash("0x27f12abfe35860a9a927b465bb3d4a9c23c8428174b83f278fe45ed7b4da2662") {
		t.Errorf("Expected topic 0x27f12abfe35860a9a927b465bb3d4a9c23c8428174b83f278fe45ed7b4da2662, got %s", logs[0].Topics[0].Hex())
	}
	if len(logs[0].Data) != 32 {
		t.Errorf("Expected data length 32, got %d", len(logs[0].Data))
	}
	hexStr := hex.EncodeToString(logs[0].Data)
	require.Equal(t, "000000000000000000000000000000000000000000000000004e7d9f51656762", hexStr)
}
