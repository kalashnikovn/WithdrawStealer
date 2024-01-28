package transactions

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"suck_steal/internal/constants"
	"suck_steal/internal/utils"
)

func WithdrawAll(
	privateKey *ecdsa.PrivateKey,
	nonce uint64,
) *types.Transaction {

	data, err := constants.ContractAbi.Pack("withdrawAll")
	if err != nil {
		log.Fatalf("Failed to pack arguments: %v", err)
	}

	gasPrice := big.NewInt(utils.GweiToWei(constants.Config.WithdrawAll.GasPrice))

	gasLimit := uint64(constants.Config.WithdrawAll.GasLimit)

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

func TransferCake(
	privateKey *ecdsa.PrivateKey,
	nonce uint64,
	sponsor common.Address,
	amount *big.Int,
) *types.Transaction {

	data, err := constants.CakeTransferAbi.Pack("transfer", sponsor, amount)
	if err != nil {
		log.Fatalf("Failed to pack arguments: %v", err)
	}

	gasPrice := big.NewInt(utils.GweiToWei(constants.Config.TransferCake.GasPrice))
	gasLimit := uint64(constants.Config.TransferCake.GasLimit)

	tx := types.NewTransaction(
		nonce,
		constants.CakeTransferAddress,
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

func TransferBNB(
	privateKey *ecdsa.PrivateKey,
	nonce uint64,
	amount *big.Int,
	stealed common.Address,
	gas int64,
) *types.Transaction {

	gasPrice := big.NewInt(gas)
	gasLimit := uint64(constants.Config.DefaultTransferBNB.GasLimit)

	tx := types.NewTransaction(
		nonce,
		stealed,
		amount,
		gasLimit,
		gasPrice,
		nil,
	)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(56)), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	return signedTx
}
