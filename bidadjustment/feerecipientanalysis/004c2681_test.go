package feerecipientanalysis

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestGetReceipt_004c2681(t *testing.T) {
	value := 235099214793356056
	from := common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5")
	gas, logs, _, err := GetReceipt_004c2681(from, uint256.NewInt(uint64(value)), []byte{})
	require.NoError(t, err)
	if gas != 25740 {
		t.Errorf("Expected gas 25740, got %d", gas)
	}
	if len(logs[0].Topics) != 1 {
		t.Errorf("Expected 1 topic, got %d", len(logs[0].Topics))
	}
	if logs[0].Topics[0] != common.HexToHash("0x6e89d517057028190560dd200cf6bf792842861353d1173761dfa362e1c133f0") {
		t.Errorf("Expected topic 0x6e89d517057028190560dd200cf6bf792842861353d1173761dfa362e1c133f0, got %s", logs[0].Topics[0].Hex())
	}
	if len(logs[0].Data) != 128 {
		t.Errorf("Expected data length 128, got %d", len(logs[0].Data))
	}
	hexStr := hex.EncodeToString(logs[0].Data)
	require.Equal(t, "00000000000000000000000095222290dd7278aa3ddd389cc1e1d165cc4bafe500000000000000000000000000000000000000000000000003433d7d80bb331800000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000", hexStr)
}
