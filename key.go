package hdwallets

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/cpacia/bchutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/iqoption/zecutil"
)

type Key struct {
	MasterKey *hdkeychain.ExtendedKey
	opt       *Option
	seed      []byte
}

// 参考 https://github.com/iancoleman/bip39
func NewKey(seed []byte, params *Option) (*Key, error) {
	k := &Key{
		seed: seed,
		opt:  params,
	}
	err := k.newMaster()
	return k, err
}

// SeedString
func (k *Key) SeedString() string {
	return hex.EncodeToString(k.seed)
}

// NewMaster 创建一个master key
func (k *Key) newMaster() (err error) {
	k.MasterKey, err = hdkeychain.NewMaster(k.seed, &k.opt.Params)
	return
}

// GetBIP32RootKeyString
func (k *Key) GetBIP32RootKeyString() string {
	return k.MasterKey.String()
}

func (k *Key) NewMasterkey() (*ExtendedKey, error) {
	c := &ExtendedKey{
		Key: k.MasterKey,
		opt: k.opt,
	}
	return c, nil
}

func (k *Key) NewChildkey(dp *DerivationPath) (*ExtendedKey, error) {

	c := &ExtendedKey{
		Key: k.MasterKey,
		opt: k.opt,
	}
	var err error

	for _, i := range dp.Path() {
		c.Key, err = c.Key.Child(i)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

type ExtendedKey struct {
	Key *hdkeychain.ExtendedKey
	opt *Option
}

func (k *ExtendedKey) Address() (string, error) {
	switch k.opt.Symbol {
	case "BCH":
		return k.AddressBCH()
	case "ZEC":
		return k.AddressZEC()
	case "ETH", "ETC":
		return k.AddressETH()
	default:
		return k.AddressBTC()
	}
}

func (k *ExtendedKey) PublicKey() (string, error) {
	return k.PublicHex(true)
}

func (k *ExtendedKey) PrivateKey() (string, error) {
	if networkIsEthereum(k.opt.Symbol) {
		return k.PrivateHex()
	}
	return k.PrivateWIF(true)

}

func (k *ExtendedKey) PublicHash() ([]byte, error) {
	address, err := k.Key.Address(&k.opt.Params)
	if err != nil {
		return nil, err
	}
	return address.ScriptAddress(), nil
}

func (k *ExtendedKey) AddressP2WPKH() (string, error) {
	pubHash, err := k.PublicHash()
	if err != nil {
		return "", err
	}

	addr, err := btcutil.NewAddressWitnessPubKeyHash(pubHash, &k.opt.Params)
	if err != nil {
		return "", err
	}

	return addr.EncodeAddress(), nil
}

// AddressP2WPKHInP2SH
func (k *ExtendedKey) AddressP2WPKHInP2SH() (string, error) {
	pubHash, err := k.PublicHash()
	if err != nil {
		return "", err
	}

	addr, err := btcutil.NewAddressWitnessPubKeyHash(pubHash, &k.opt.Params)
	if err != nil {
		return "", err
	}

	script, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return "", err
	}

	addr1, err := btcutil.NewAddressScriptHash(script, &k.opt.Params)
	if err != nil {
		return "", err
	}

	return addr1.EncodeAddress(), nil
}

func (k *ExtendedKey) AddressBTC() (string, error) {

	address, err := k.Key.Address(&k.opt.Params)
	if err != nil {
		return "", err
	}
	return address.EncodeAddress(), nil
}

func (k *ExtendedKey) AddressETH() (string, error) {
	privateECDSA, err := k.Key.ECPrivKey()
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(privateECDSA.ToECDSA().PublicKey).Hex(), nil
}

func (k *ExtendedKey) AddressBCH() (string, error) {
	address, err := k.Key.Address(&k.opt.Params)
	if err != nil {
		return "", err
	}
	addr, err := bchutil.NewCashAddressPubKeyHash(address.ScriptAddress(), &k.opt.Params)
	if err != nil {
		return "", err
	}
	data := addr.EncodeAddress()
	return "bitcoincash:" + data, nil
}

func (k *ExtendedKey) AddressZEC() (string, error) {
	privatekey, err := k.Key.ECPrivKey()
	if err != nil {
		return "", err
	}
	wif, err := btcutil.NewWIF(privatekey, &k.opt.Params, true)
	if err != nil {
		return "", err
	}
	return zecutil.Encode(wif.PrivKey.PubKey().SerializeCompressed(), &k.opt.Params)
}

func (k *ExtendedKey) PublicHex(compress bool) (string, error) {
	ecpubKey, err := k.Key.ECPubKey()
	if err != nil {
		return "", err
	}
	if compress {
		return hex.EncodeToString(ecpubKey.SerializeCompressed()), nil
	}

	return hex.EncodeToString(ecpubKey.SerializeUncompressed()), nil
}

func (k *ExtendedKey) PrivateHex() (string, error) {
	ecprivKey, err := k.Key.ECPrivKey()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ecprivKey.Serialize()), nil
}

// PrivateWIF
func (k *ExtendedKey) PrivateWIF(compress bool) (string, error) {
	ecprivKey, err := k.Key.ECPrivKey()
	if err != nil {
		return "", err
	}
	wif, err := btcutil.NewWIF(ecprivKey, &k.opt.Params, compress)
	if err != nil {
		return "", err
	}

	return wif.String(), nil
}
