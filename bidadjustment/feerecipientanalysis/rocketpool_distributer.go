package feerecipientanalysis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	EOA_RocketpoolDistributer = common.FromHex("608060405236600a57005b60008054604080517f21f8a7210000000000000000000000000000000000000000000000000000000081527ffb3483c04a4c870b23f9315ce407aceff6ee8dd58eb34aa5becfb781351d3fb86004820152905173ffffffffffffffffffffffffffffffffffffffff909216916321f8a72191602480820192602092909190829003018186803b158015609b57600080fd5b505afa15801560ae573d6000803e3d6000fd5b505050506040513d602081101560c357600080fd5b505190503660008037600080366000845af43d6000803e80801560e5573d6000f35b3d6000fdfea2646970667358221220eff696097e25305c6d6b19a8a6eebb3c8b47de4d235eade3a84bd4f50c48119664736f6c63430007060033")
)

func GetReceipt_RocketpoolDistributer() (uint64, []*types.Log, []*types.Log) {
	//Example https://etherscan.io/address/0xd4E96eF8eee8678dBFf4d535E033Ed1a4F7605b7
	// https://etherscan.io/tx/0xee2c5d0f6dd84eadf90340356e69c0cdb3149d9983cc7c73827c272822ee4953
	return 21000 + 33, []*types.Log{}, []*types.Log{}
}

// label_0000:
// 	// Inputs[1] { @0005  msg.data.length }
// 	3 0000    60  PUSH1 0x80
//  3 0002    60  PUSH1 0x40
// 	12 0004    52  MSTORE
// 	2 0005    36  CALLDATASIZE
// 	3 0006    60  PUSH1 0x0a
// 	10 0008    57  *JUMPI
// 	// Stack delta = +0
// 	// Outputs[1] { @0004  memory[0x40:0x60] = 0x80 }
// 	// Block ends with conditional jump to 0x000a, if msg.data.length

// label_0009:
// 	// Incoming jump from 0x0008, if not msg.data.length
// 	0009    00  *STOP
// 	// Stack delta = +0
// 	// Outputs[1] { @0009  stop(); }
// 	// Block terminates

// = 3+3+12+2+3+10 = 33
