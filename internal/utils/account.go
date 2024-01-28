package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  common.Address
}

func NewAccount(pk string) Account {
	accPrivateKey, accAddress := StrToWallet(pk)

	return Account{
		PrivateKey: accPrivateKey,
		PublicKey:  accAddress,
	}

}
