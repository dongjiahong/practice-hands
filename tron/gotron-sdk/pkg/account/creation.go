package account

import (
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/fbsobreira/gotron-sdk/pkg/mnemonic"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
)

// Creation struct for account
type Creation struct {
	Name               string
	Passphrase         string
	Mnemonic           string
	MnemonicPassphrase string
	HdAccountNumber    *uint32
	HdIndexNumber      *uint32
}

// New create new name
func New() string {
	return "New Account"
}

// IsValidPassphrase check if strong
func IsValidPassphrase(pass string) bool {
	// TODO: force strong password
	return true
}

// CreateNewLocalAccount assumes all the inputs are valid, legitmate
func CreateNewLocalAccount(candidate *Creation) error {
	ks := store.FromAccountName(candidate.Name) // 获取目录: /home/lele/.tronctl/account-keys/vimer
	if candidate.Mnemonic == "" {
		candidate.Mnemonic = mnemonic.Generate() // 生成助24个记词
	}
	// Hardcoded index of 0 for brandnew account., => 新账号的硬编码序号为0
	private, _ := keys.FromMnemonicSeedAndPassphrase(candidate.Mnemonic, candidate.MnemonicPassphrase, 0) // 返回一个<私钥，公钥>密钥对
	_, err := ks.ImportECDSA(private.ToECDSA(), candidate.Passphrase)                                     // 生成密钥文件UTC--2021-03-18T--2o93u42o3u02n2i382n293090129en0
	if err != nil {
		return err
	}
	return nil
}
