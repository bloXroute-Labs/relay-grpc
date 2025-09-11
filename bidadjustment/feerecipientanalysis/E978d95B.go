package feerecipientanalysis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

// 0xE978d95B437D75826ABa2ED1A0BDb534f173E28c
// https://etherscan.io/tx/0x958f8313bbe2b77d99f62d387a1d2bc010d346d87b7e0821ab0d08a987039665#eventlog
// https://ethervm.io/decompile/0xE978d95B437D75826ABa2ED1A0BDb534f173E28c#011a

var (
	FeeRecipient_E978d95B = common.HexToAddress("0xE978d95B437D75826ABa2ED1A0BDb534f173E28c")
)

func GetReceipt_E978d95B(newBalance *uint256.Int) (uint64, []*types.Log, []*types.Log, error) {
	uint256Bytes := newBalance.Bytes32()
	//uint256Value in a 32byte array
	return 21000 + 1218, []*types.Log{{
		Address: common.HexToAddress("0xE978d95B437D75826ABa2ED1A0BDb534f173E28c"),
		Topics:  []common.Hash{common.HexToHash("0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936")},
		Data:    uint256Bytes[:],
	}}, []*types.Log{}, nil
}

// label_0000:
// 	// Inputs[1] { @0007  msg.data.length }
// 	0000    60  PUSH1 0x80 3
// 	0002    60  PUSH1 0x40 3
// 	0004    52  MSTORE 12  m[0x40:0x60] = 0x80
// 	0005    60  PUSH1 0x04 3
// 	0007    36  CALLDATASIZE 2
// 	0008    10  LT 3
// 	0009    61  PUSH2 0x00c7 3
// 	000C    57  *JUMPI  10

// 	// 3 + 3 + 12 + 3 + 2 + 3 + 3 + 10 = 39

// label_00C7:
// 	00C7    5B  JUMPDEST 1
// 	00C8    36  CALLDATASIZE 2
// 	00C9    15  ISZERO 3
// 	00CA    61  PUSH2 0x011a 3
// 	00CD    57  *JUMPI 10

// 	// 1 + 2 + 3 + 3 + 10 = 19

// 	label_011A:
// 	// Incoming jump from 0x00CD, if !msg.data.length
// 	// Inputs[5]
// 	// {
// 	//     @011E  memory[0x40:0x60]
// 	//     @011F  address(this)
// 	//     @0120  address(this).balance
// 	//     @0124  memory[0x40:0x60]
// 	//     @014E  memory[memory[0x40:0x60]:memory[0x40:0x60] + 0x20 + (memory[0x40:0x60] - memory[0x40:0x60])]
// 	// }
// 	011A    5B  JUMPDEST 1
// 	011B    60  PUSH1 0x40 3
// 	011D    80  DUP1 3 0x40,0x40
// 	011E    51  MLOAD 3 (no mem increase) m[0x40:0x60] = 0x80, 0x40
// 	011F    30  ADDRESS 2 a,0x80,0x40
// 	0120    31  BALANCE 100 (warm address) b,0x80,0x40
// 	0121    81  DUP2 3 0x80,b,0x80,0x40
// 	0122    52  MSTORE 9 [0x80, b] -> 3 + 6 = 9 [0x80,0x40] m[0x80:0xa0] = b 0x80,0x40
// 	0123    90  SWAP1 3 0x40,0x80
// 	0124    51  MLOAD 3 0x80,0x80
// 	0125    7F  PUSH32 0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936 3 0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936,0x80,0x80
// 	0146    91  SWAP2 3 // 0x80,0x80,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	0147    81  DUP2 3 // 0x80,0x80,0x80,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	0148    90  SWAP1 3 // 0x80,0x80,0x80,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	0149    03  SUB 3 // 0,0x80,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	014A    60  PUSH1 0x20 3 0x20,0,0x80,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	014C    01  ADD 3 0x20,0x80,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	014D    90  SWAP1 3 // 0x80,0x20,0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 	014E    A1  LOG1 ... gas_cost = 375 + 375 * num_topics + 8 * data_size + mem_expansion_cost
// 				m[0x80:0xa0], 0xc4c14883ae9fd8e26d5d59e3485ed29fd126d781d7e498a4ca5c54c8268e4936
// 				// 375 + 375*1 + 8*32 + 0 = 1006
// 	014F    00  *STOP 0

// // 1 + 3 + 3 + 3 + 2 + 100 + 3 + 9 + 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 + 3 = 154
// // 154 + 375 + 375 + 8 * 32 + 0 = 1160
// // 1160 + 39 + 19 = 1218
