package feerecipientanalysis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

// 0x004c26816dA9219CF3408e84eDD9716Df4B5739A - LIDO

// https://etherscan.io/address/0x004c26816dA9219CF3408e84eDD9716Df4B5739A#code
// https://etherscan.io/tx/0x33114da82089a06c391aae7b6376cee38494dba6ee2a347942de668f95bf15ca
// https://ethervm.io/decompile/0x004c26816dA9219CF3408e84eDD9716Df4B5739A

var (
	FeeRecipient_004c2681 = common.HexToAddress("0x004c26816dA9219CF3408e84eDD9716Df4B5739A")
)

func GetReceipt_004c2681(from common.Address, value *uint256.Int, msg []byte) (uint64, []*types.Log, []*types.Log, error) {
	data := make([]byte, 128)
	copy(data[12:], from.Bytes())
	vBytes := value.Bytes32()
	copy(data[32:], vBytes[:])
	data[95] = 0x60
	return 21000 + 4740, []*types.Log{{
		Address: common.HexToAddress("0x004c26816dA9219CF3408e84eDD9716Df4B5739A"),
		Topics:  []common.Hash{common.HexToHash("0x6e89d517057028190560dd200cf6bf792842861353d1173761dfa362e1c133f0")},
		Data:    data,
	}}, []*types.Log{}, nil
}

// 2 0000    36  CALLDATASIZE
// 2 0001    3D  RETURNDATASIZE
// 2 0002    3D  RETURNDATASIZE
// 2 0003    37  CALLDATACOPY
// 2 0004    3D  RETURNDATASIZE
// 2 0005    3D  RETURNDATASIZE
// 2 0006    3D  RETURNDATASIZE
// 3 0007    36  CALLDATASIZE
// 2 0008    3D  RETURNDATASIZE
// 3 0009    73  PUSH20 0xe8e847cf573fc8ed75621660a36affd18c543d7e
// 2 001E    5A  GAS
// 2600 + X 001F    F4  DELEGATECALL
// 2 0020    3D  RETURNDATASIZE
// 3 0021    82  DUP3
// 3 0022    80  DUP1
// 3 0023    3E  RETURNDATACOPY
// 3 0024    90  SWAP1
// 2 0025    3D  RETURNDATASIZE
// 3 0026    91  SWAP2
// 3 0027    60  PUSH1 0x2b
// 10 0029    57  *JUMPI

// 1 002B    5B  JUMPDEST
// 002C    F3  *RETURN
// Stack delta = -2
// Outputs[1] { @002C  return memory[stack[-1]:stack[-1] + stack[-2]]; }
// Block terminates

//2*7 + 3+2+3+2+X+2+3*4+2+3*2+10+1 = 2657 + X
//= 2657 + 2083
//=4740

// 0xe8e847cf573fc8ed75621660a36affd18c543d7e
// https://etherscan.io/address/0xe8e847cf573fc8ed75621660a36affd18c543d7e#code

// label_0000:
// 	// Inputs[1] { @0007  msg.data.length }
// 	3 0000    60  PUSH1 0x80
// 	3 0002    60  PUSH1 0x40
// 	12 0004    52  MSTORE
// 	3 0005    60  PUSH1 0x04
// 	2 0007    36  CALLDATASIZE
// 	3 0008    10  LT
// 	3 0009    61  PUSH2 0x00a0
// 	10 000C    57  *JUMPI
// 	// Stack delta = +0
// 	// Outputs[1] { @0004  memory[0x40:0x60] = 0x80 }
// 	// Block ends with conditional jump to 0x00a0, if msg.data.length < 0x04
// 	//3+3+12+3+2+3+3+10 = 39

// 	label_00A0:
// 	// Incoming jump from 0x000C, if msg.data.length < 0x04
// 	// Inputs[1] { @00A1  msg.data.length }
// 	1 00A0    5B  JUMPDEST
// 	2 00A1    36  CALLDATASIZE
// 	3 00A2    61  PUSH2 0x013b
// 	10 00A5    57  *JUMPI
// 	// Stack delta = +0
// 	// Block ends with conditional jump to 0x013b, if msg.data.length
//     //1+2+3+10 = 16

// 	label_00A6:
// 	// Incoming jump from 0x00A5, if not msg.data.length
// 	// Inputs[1] { @00A8  msg.value }
// 	3 00A6    60  PUSH1 0x00
// 	2 00A8    34  CALLVALUE
// 	3 00A9    11  GT
// 	3 00AA    15  ISZERO
// 	3 00AB    61  PUSH2 0x0139
// 	10 00AE    57  *JUMPI
// 	// 3+2+3+3+3+10 = 24

// 	// Stack delta = +0
// 	// Block ends with conditional jump to 0x0139, if !(msg.value > 0x00)

// 	label_00AF:
// 	// Incoming jump from 0x00AE, if not !(msg.value > 0x00)
// 	// Inputs[7]
// 	// {
// 	//     @00D0  msg.sender
// 	//     @00D1  msg.value
// 	//     @00D4  msg.data.length
// 	//     @00D7  memory[0x40:0x60]
// 	//     @0112  msg.data[0x00:0x00 + msg.data.length]
// 	//     @0133  memory[0x40:0x60]
// 	//     @0138  memory[memory[0x40:0x60]:memory[0x40:0x60] + (0x20 + 0x20 + 0x20 + 0x20 + memory[0x40:0x60] + (msg.data.length + 0x1f & ~0x1f)) - memory[0x40:0x60]]
// 	// }
// 	3 00AF    7F  PUSH32 0x6e89d517057028190560dd200cf6bf792842861353d1173761dfa362e1c133f0
// 	2 00D0    33  CALLER
// 	2 00D1    34  CALLVALUE
// 	3 00D2    60  PUSH1 0x00
// 	2 00D4    36  CALLDATASIZE
// 	3 00D5    60  PUSH1 0x40
// 	3 00D7    51  MLOAD [0x40:0x60]
// 	3 00D8    80  DUP1
// 	3 00D9    85  DUP6
// 	3 00DA    73  PUSH20 0xffffffffffffffffffffffffffffffffffffffff
// 	3 00EF    16  AND
// 	3 00F0    81  DUP2
// 	9 00F1    52  MSTORE [0x80:0xa0]
// 	3 00F2    60  PUSH1 0x20
// 	3 00F4    01  ADD
// 	3 00F5    84  DUP5
// 	3 00F6    81  DUP2
// 	6 00F7    52  MSTORE [0xa0:0xc0]
// 	3 00F8    60  PUSH1 0x20
// 	3 00FA    01  ADD
// 	3 00FB    80  DUP1
// 	3 00FC    60  PUSH1 0x20
// 	3 00FE    01  ADD
// 	3 00FF    82  DUP3
// 	3 0100    81  DUP2
// 	3 0101    03  SUB
// 	3 0102    82  DUP3
// 	6 0103    52  MSTORE
// 	3 0104    84  DUP5
// 	3 0105    84  DUP5
// 	3 0106    82  DUP3
// 	3 0107    81  DUP2
// 	3 0108    81  DUP2
// 	6 0109    52  MSTORE
// 	3 010A    60  PUSH1 0x20
// 	3 010C    01  ADD
// 	3 010D    92  SWAP3
// 	2 010E    50  POP
// 	3 010F    80  DUP1
// 	3 0110    82  DUP3
// 	3 0111    84  DUP5
// 	3 0112    37  CALLDATACOPY
// 	3 0113    60  PUSH1 0x00
// 	3 0115    81  DUP2
// 	3 0116    84  DUP5
// 	3 0117    01  ADD
// 	6 0118    52  MSTORE
// 	3 0119    60  PUSH1 0x1f
// 	3 011B    19  NOT
// 	3 011C    60  PUSH1 0x1f
// 	3 011E    82  DUP3
// 	3 011F    01  ADD
// 	3 0120    16  AND
// 	3 0121    90  SWAP1
// 	2 0122    50  POP
// 	3 0123    80  DUP1
// 	3 0124    83  DUP4
// 	3 0125    01  ADD
// 	3 0126    92  SWAP3
// 	2 0127    50  POP
// 	2 0128    50  POP
// 	2 0129    50  POP
// 	3 012A    95  SWAP6
// 	2 012B    50  POP
// 	2 012C    50  POP
// 	2 012D    50  POP
// 	2 012E    50  POP
// 	2 012F    50  POP
// 	2 0130    50  POP
// 	3 0131    60  PUSH1 0x40
// 	3 0133    51  MLOAD
// 	3 0134    80  DUP1
// 	3 0135    91  SWAP2
// 	3 0136    03  SUB
// 	3 0137    90  SWAP1
// 	375+375*1+8*(32+32+32+32)+0=1774 0138    A1  LOG1
// 	1 0139    5B  JUMPDEST
// 	013A    00  *STOP

// 	= 3+2+2+3+2+3+3+3+3+3+3+3+9+3+3+3+3+6+3+3+3+3+3+3+3+3+3+6+3+3+3+3+3+6+3+3+3+2+3+3+3+3+3+3+3+3+6+3+3+3+3+3+3+3+2+3+3+3+3+2+2+2+3+2+2+2+2+2+2+3+3+3+3+3+3+1+1774
// 	= 2004

// 	total = 2004 + 24+ 16 + 39 = 2083
