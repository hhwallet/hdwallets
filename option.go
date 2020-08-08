package hdwallets

import (
	"github.com/btcsuite/btcd/chaincfg"
)

type Option struct {
	Symbol string
	chaincfg.Params
}

var (
	BTCParams = Option{
		Symbol: "BTC",
		Params: chaincfg.MainNetParams,
	}
	BTCTestnetParams = Option{
		Symbol: "BTCTest",
		Params: chaincfg.TestNet3Params,
	}

	LTCParams = Option{
		Symbol: "LTC",
		Params: chaincfg.MainNetParams,
	}
	DOGEParams = Option{
		Symbol: "DOGE",
		Params: chaincfg.MainNetParams,
	}
	DASHParams = Option{
		Symbol: "DASH",
		Params: chaincfg.MainNetParams,
	}
	BCHParams = Option{
		Symbol: "BCH",
		Params: chaincfg.MainNetParams,
	}
	ETHParams = Option{
		Symbol: "ETH",
		Params: chaincfg.MainNetParams,
	}
	ETCParams = Option{
		Symbol: "ETC",
		Params: chaincfg.MainNetParams,
	}
	ZECParams = Option{
		Symbol: "ZEC",
		Params: chaincfg.MainNetParams,
	}
)

// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
func init() {
	LTCParams.Bech32HRPSegwit = "ltc"
	LTCParams.PubKeyHashAddrID = 0x30 // 48
	LTCParams.ScriptHashAddrID = 0x32 // 50
	LTCParams.PrivateKeyID = 0xB0     // 176
	LTCParams.HDCoinType = 2

	ZECParams.HDCoinType = 133

	DOGEParams.PubKeyHashAddrID = 0x1e // 30
	DOGEParams.ScriptHashAddrID = 0x16 // 22
	DOGEParams.PrivateKeyID = 0x9e     // 158
	DOGEParams.HDCoinType = 3

	DASHParams.PubKeyHashAddrID = 0x4c // 76
	DASHParams.ScriptHashAddrID = 0x10 // 16
	DASHParams.PrivateKeyID = 0xcc     // 204
	DASHParams.HDCoinType = 5

	BCHParams.PubKeyHashAddrID = 0x00 // 0
	BCHParams.ScriptHashAddrID = 0x05 // 5
	BCHParams.PrivateKeyID = 0x80     // 128
	BCHParams.HDCoinType = 145

	ETHParams.HDCoinType = 60

	ETCParams.HDCoinType = 61

}

func networkIsEthereum(symbol string) bool {
	return symbol == "ETH" || symbol == "ETC" ||
		symbol == "EWT" || symbol == "PIRL" ||
		symbol == "MIX" || symbol == "MOAC" ||
		symbol == "MUSIC" || symbol == "POA" ||
		symbol == "EXP" || symbol == "CLO" ||
		symbol == "DXN" || symbol == "ELLA" ||
		symbol == "ESN" || symbol == "VET" ||
		symbol == "ERE"
}
