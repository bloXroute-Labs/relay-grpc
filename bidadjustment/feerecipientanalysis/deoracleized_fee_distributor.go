package feerecipientanalysis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	EOA_DeoracleizedFeeRecipient = common.FromHex("0x363d3d373d3d3d363d733fcd8d9acac042095dfba53f4c40c74d19e2e9d95af43d82803e903d91602b57fd5bf3")
)

func GetReceipt_DeoracleizedFeeRecipient() (uint64, []*types.Log, []*types.Log) {
	//Example https://etherscan.io/address/0x8C07b79C79911f87e70cD9E549c10145e891Ae6d
	return 25868, []*types.Log{}, []*types.Log{}
}

// label_0000:
// 	// Inputs[15]
// 	// {
// 	//     @0000  msg.data.length
// 	//     @0001  returndata.length
// 	//     @0002  returndata.length
// 	//     @0003  msg.data[returndata.length:returndata.length + msg.data.length]
// 	//     @0004  returndata.length
// 	//     @0005  returndata.length
// 	//     @0006  returndata.length
// 	//     @0007  msg.data.length
// 	//     @0008  returndata.length
// 	//     @001E  msg.gas
// 	//     @001F  address(0x3fcd8d9acac042095dfba53f4c40c74d19e2e9d9).delegatecall.gas(msg.gas)(memory[returndata.length:returndata.length + msg.data.length])
// 	//     @001F  memory[returndata.length:returndata.length + msg.data.length]
// 	//     @0020  returndata.length
// 	//     @0023  returndata[returndata.length:returndata.length + returndata.length]
// 	//     @0025  returndata.length
// 	// }
// 	0000    36  CALLDATASIZE
// 	0001    3D  RETURNDATASIZE
// 	0002    3D  RETURNDATASIZE
// 	0003    37  CALLDATACOPY
// 	0004    3D  RETURNDATASIZE
// 	0005    3D  RETURNDATASIZE
// 	0006    3D  RETURNDATASIZE
// 	0007    36  CALLDATASIZE
// 	0008    3D  RETURNDATASIZE
// 	0009    73  PUSH20 0x3fcd8d9acac042095dfba53f4c40c74d19e2e9d9
// 	001E    5A  GAS
// 	001F    F4  DELEGATECALL
// 	0020    3D  RETURNDATASIZE
// 	0021    82  DUP3
// 	0022    80  DUP1
// 	0023    3E  RETURNDATACOPY
// 	0024    90  SWAP1
// 	0025    3D  RETURNDATASIZE
// 	0026    91  SWAP2
// 	0027    60  PUSH1 0x2b
// 	0029    57  *JUMPI
// 	// Stack delta = +2
// 	// Outputs[5]
// 	// {
// 	//     @0003  memory[returndata.length:returndata.length + msg.data.length] = msg.data[returndata.length:returndata.length + msg.data.length]
// 	//     @001F  memory[returndata.length:returndata.length + returndata.length] = address(0x3fcd8d9acac042095dfba53f4c40c74d19e2e9d9).delegatecall.gas(msg.gas)(memory[returndata.length:returndata.length + msg.data.length])
// 	//     @0023  memory[returndata.length:returndata.length + returndata.length] = returndata[returndata.length:returndata.length + returndata.length]
// 	//     @0024  stack[1] = returndata.length
// 	//     @0026  stack[0] = returndata.length
// 	// }
// 	// Block ends with conditional jump to 0x002b, if address(0x3fcd8d9acac042095dfba53f4c40c74d19e2e9d9).delegatecall.gas(msg.gas)(memory[returndata.length:returndata.length + msg.data.length])

// label_002A:

// label_0000:
// 	// Inputs[1] { @0007  msg.data.length }
// 	0000    60  PUSH1 0x80
// 	0002    60  PUSH1 0x40
// 	0004    52  MSTORE
// 	0005    60  PUSH1 0x04
// 	0007    36  CALLDATASIZE
// 	0008    10  LT
// 	0009    15  ISZERO
// 	000A    61  PUSH2 0x0022
// 	000D    57  *JUMPI
// 	// Stack delta = +0
// 	// Outputs[1] { @0004  memory[0x40:0x60] = 0x80 }
// 	// Block ends with conditional jump to 0x0022, if !(msg.data.length < 0x04)

// label_000E:
// 	// Incoming jump from 0x0123, if 0xf818e093 - stack[-1]
// 	// Incoming jump from 0x000D, if not !(msg.data.length < 0x04)
// 	// Inputs[1] { @000F  msg.data.length }
// 	000E    5B  JUMPDEST
// 	000F    36  CALLDATASIZE
// 	0010    15  ISZERO
// 	0011    61  PUSH2 0x0018
// 	0014    57  *JUMPI

// label_0018:
// // Incoming jump from 0x0014, if !msg.data.length
// 0018    5B  JUMPDEST
// 0019    61  PUSH2 0x0020
// 001C    61  PUSH2 0x142f
// 001F    56  *JUMP
// // Stack delta = +1
// // Outputs[1] { @0019  stack[0] = 0x0020 }
// // Block ends with call to 0x142f, returns to 0x0020
// label_142F:
// 	// Incoming call from 0x001F, returns to 0x0020
// 	// Inputs[1] { @1432  storage[0x04] }
// 	142F    5B  JUMPDEST
// 	1430    60  PUSH1 0x04
// 	1432    54  SLOAD
// 	1433    60  PUSH1 0x60
// 	1435    1C  SHR
// 	1436    15  ISZERO
// 	1437    61  PUSH2 0x143c
// 	143A    57  *JUMPI
// 	// Stack delta = +0
// 	// Block ends with conditional jump to 0x143c, if !(storage[0x04] >> 0x60)

// label_143B:
// 	// Incoming jump from 0x143A, if not !(storage[0x04] >> 0x60)
// 	// Inputs[1] { @143B  stack[-1] }
// 	143B    56  *JUMP
// 	// Stack delta = -1
// 	// Block ends with unconditional jump to stack[-1]

// 	label_0020:
// 	// Incoming return from call to 0x03B8 at 0x09B2
// 	// Incoming jump from 0x06DB
// 	// Incoming return from call to 0x142F at 0x001F
// 	// Incoming return from call to 0x144E at 0x0250
// 	0020    5B  JUMPDEST
// 	0021    00  *STOP
// 	// Stack delta = +0
// 	// Outputs[1] { @0021  stop(); }
// 	// Block terminates
