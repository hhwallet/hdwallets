package hdwallets

import (
	"fmt"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// m / purpose' / coin_type' / account' / change / address_index
type DerivationPath struct {
	Purpose uint32
	CoinType uint32
	Account uint32
	Change uint32
	AddressIndex uint32
}

func NewBIP44DerivationPath(coinType,account,change,addressIndex uint32) *DerivationPath{

	return NewDerivationPath(44,coinType,account,change,addressIndex)
}

func NewDerivationPath(purpose,coinType,account,change,addressIndex uint32) *DerivationPath{
		return &DerivationPath{
		Purpose:      purpose + 	hdkeychain.HardenedKeyStart,
		CoinType:     coinType + hdkeychain.HardenedKeyStart,
		Account:      account + hdkeychain.HardenedKeyStart,
		Change:       change,
		AddressIndex: addressIndex,
	}
}

func (d *DerivationPath) PathString() string{
	return fmt.Sprintf("m/%d'/%d'/%d'/%d/%d",d.Purpose-hdkeychain.HardenedKeyStart,d.CoinType-hdkeychain.HardenedKeyStart,d.Account-hdkeychain.HardenedKeyStart,d.Change,d.AddressIndex)
}

func (d *DerivationPath) Path() []uint32{
	return []uint32{
		d.Purpose,
		d.CoinType,
		d.Account,
		d.Change,
		d.AddressIndex,
	}
}
