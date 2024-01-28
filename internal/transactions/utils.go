package transactions

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetGasEstimateFromTransaction(tx *types.Transaction, client *ethclient.Client, account common.Address) (uint64, error) {
	ctx := context.Background()
	callMsg := ethereum.CallMsg{
		From:     account,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	gasEstimate, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return 0, err
	}
	return gasEstimate, nil
}
