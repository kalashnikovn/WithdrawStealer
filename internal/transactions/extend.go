package transactions

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"suck_steal/internal/constants"
	"suck_steal/internal/utils"
)

func Extend(
	privateKey *ecdsa.PrivateKey,
	nonce uint64,
) *types.Transaction {

	data, err := constants.ContractAbi.Pack("deposit", big.NewInt(0), big.NewInt(constants.Config.Deposit.ExtendTime))
	if err != nil {
		log.Fatalf("Failed to pack arguments: %v", err)
	}

	gasPrice := big.NewInt(utils.GweiToWei(constants.Config.Deposit.GasPrice))

	gasLimit := uint64(constants.Config.Deposit.GasLimit)

	tx := types.NewTransaction(
		nonce,
		constants.ContractAddress,
		big.NewInt(0),
		gasLimit,
		gasPrice,
		data,
	)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(56)), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	return signedTx
}
