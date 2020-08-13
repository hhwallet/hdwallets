package hdwallets

import (
	"fmt"
	"testing"
)

func TestKey_Address(t *testing.T) {
	var lang = "english"
	mn,_ := NewMnemonic(12,lang)
	//fmt.Println(mn)
	mn = "try actor open dwarf cushion admit trend practice mammal income danger duck"
	fmt.Println("mnemonic:", mn)
	seed, _ := NewSeed(mn, "", lang)
	params := &XRPParams
	// ltc: LcamvNe5XZEorxsU8PxkqqTsaX1aByCA96 037e69f608e9a975748014b5fa0b66766889365b52304ab3fce5242675617e573f TApxsqFh5AgcAMkeu19LQb9nfhUryJj8zKomQvkQeMpQ2rA3DhH2
	k, _ := NewKey(seed, params)
	fmt.Println("seed", k.SeedString())
	path := NewBIP44DerivationPath(params.HDCoinType, 0, 0, 0)
	fmt.Println("path:", path.PathString())
	ck, _ := k.NewChildkey(path)
	fmt.Print("address:")
	fmt.Println(ck.Address())
	fmt.Print("publicKey:")
	fmt.Println(ck.PublicKey())
	fmt.Print("privateKey:")
	fmt.Println(ck.PrivateKey())
}
