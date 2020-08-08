# hdwallets

生成bip44地址
当前支持的币种：
- BTC
- LTC
- DOGE
- DASH
- ETH
- ETC
- BCH
- ZEC

## Example
```
    var lang = "english"
	mn,_ := NewMnemonic(12,lang)
	//fmt.Println(mn)
	fmt.Println("mnemonic:", mn)
	seed, _ := NewSeed(mn, "", lang)
	params := &BTCParams
	k, _ := NewKey(seed, params)
	fmt.Println("seed:", k.SeedString())
	path := NewBIP44DerivationPath(params.HDCoinType, 0, 0, 0)
	fmt.Println("path:", path.PathString())
	ck, _ := k.NewChildkey(path)
	fmt.Print("address:")
	fmt.Println(ck.Address())
	fmt.Print("publicKey:")
	fmt.Println(ck.PublicKey())
	fmt.Print("privateKey:")
	fmt.Println(ck.PrivateKey())

    // result 
    // mnemonic: transfer purity bright above course dial forum attend bacon tunnel unit dawn
    // seed bd0548c7ead8be76dbc48af9f1a03513caadeede8f1a6c46bda1813027d067b8f8d9cd7a01f7bc9d537bf36246b6d28e1ee2d25711fab3f2d3e30e2cb46f26ba
    // path: m/44'/0'/0'/0/0
    // address:15vKtr9NsBwcu8is5yNDmN1ogvL7VUEvjQ <nil>
    // publicKey:025567eae6131bf3695639bd333c6ea43936fc35c9c437dc337b4e7b4d173c4815 <nil>
    // privateKey:L5Bn56YD5MTJVgTeX5QZ1mJ3fgipgs2yPdrakSNbqjikP2cS9JXA <nil>
```