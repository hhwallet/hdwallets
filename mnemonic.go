package hdwallets

import (
	"errors"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

const (
	English            = "english"
	ChineseSimplified  = "chinese_simplified"
	ChineseTraditional = "chinese_traditional"
	French = "french"
	Italian = "ttalian"
	Japanese = "japanese"
	Korean = "korean"
	Spanish = "spanish"
)

func setLanguage(language string) {
	switch language {
	case ChineseSimplified:
		bip39.SetWordList(wordlists.ChineseSimplified)
	case ChineseTraditional:
		bip39.SetWordList(wordlists.ChineseTraditional)
	case French:
		bip39.SetWordList(wordlists.French)
	case Italian:
		bip39.SetWordList(wordlists.Italian)
	case Japanese:
		bip39.SetWordList(wordlists.Japanese)
	case Korean:
		bip39.SetWordList(wordlists.Korean)
	case Spanish:
		bip39.SetWordList(wordlists.Spanish)
	case English:
	default:


	}
}

// NewMnemonic 创建一个随机助记词,语言为空默认为english
func NewMnemonic(length int, language string) (string, error) {
	setLanguage(language)

	if length < 12 || length > 24 || length % 3 > 0 {
		return "",errors.New("length长度在12-24之间且是3的倍数")
	}

	entropy, err := bip39.NewEntropy(length / 3 * 32)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(entropy)
}

// NewSeed 创建一个散装的种子,语言为空默认为english
func NewSeed(mnemonic, password, language string) ([]byte, error) {
	setLanguage(language)
	return bip39.NewSeedWithErrorChecking(mnemonic, password)
}
