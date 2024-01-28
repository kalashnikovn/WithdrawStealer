package transactions

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"suck_steal/internal/constants"
	"suck_steal/internal/utils"
)

func WithdrawFullCake(sponsor utils.Account, stealer utils.Account, sponsorNonce uint64) []hexutil.Bytes {
	//sponsorNonce := utils.GetNonce(constants.Client.General, sponsor.PublicKey)
	stealerNonce := utils.GetNonce(constants.Client.General, stealer.PublicKey)

	amount := new(big.Int)
	amount.SetString(constants.Config.CakeAmount, 10)

	fee := new(big.Int)
	fee.SetString(constants.Config.Fee, 10)

	transferBnbTx := TransferBNB(sponsor.PrivateKey, sponsorNonce, fee, stealer.PublicKey, utils.GweiToWei(constants.Config.FastTransferBNB.GasPrice))
	withdrawAllTx := WithdrawAll(stealer.PrivateKey, stealerNonce)
	transferCakeTx := TransferCake(stealer.PrivateKey, stealerNonce+1, sponsor.PublicKey, amount)

	rawTxBytes1, _ := rlp.EncodeToBytes(transferBnbTx)
	rawTxBytes2, _ := rlp.EncodeToBytes(withdrawAllTx)
	rawTxBytes3, _ := rlp.EncodeToBytes(transferCakeTx)

	var rawTxs = []hexutil.Bytes{rawTxBytes1, rawTxBytes2, rawTxBytes3}

	return rawTxs
}

func WithdrawFullCakePlusNonce(sponsor utils.Account, stealer utils.Account, stealerNonce uint64, sponsorNonce uint64) []hexutil.Bytes {
	//sponsorNonce := utils.GetNonce(constants.Client.General, sponsor.PublicKey)
	//stealerNonce := utils.GetNonce(constants.Client.General, stealer.PublicKey)

	amount := new(big.Int)
	amount.SetString(constants.Config.CakeAmount, 10)

	fee := new(big.Int)
	fee.SetString(constants.Config.Fee, 10)

	transferBnbTx := TransferBNB(sponsor.PrivateKey, sponsorNonce, fee, stealer.PublicKey, utils.GweiToWei(constants.Config.FastTransferBNB.GasPrice))
	withdrawAllTx := WithdrawAll(stealer.PrivateKey, stealerNonce)
	transferCakeTx := TransferCake(stealer.PrivateKey, stealerNonce+1, sponsor.PublicKey, amount)

	rawTxBytes1, _ := rlp.EncodeToBytes(transferBnbTx)
	rawTxBytes2, _ := rlp.EncodeToBytes(withdrawAllTx)
	rawTxBytes3, _ := rlp.EncodeToBytes(transferCakeTx)

	var rawTxs = []hexutil.Bytes{rawTxBytes1, rawTxBytes2, rawTxBytes3}

	return rawTxs
}

func ExtendTime(sponsor utils.Account, stealer utils.Account) []hexutil.Bytes {
	sponsorNonce := utils.GetNonce(constants.Client.General, sponsor.PublicKey)
	stealerNonce := utils.GetNonce(constants.Client.General, stealer.PublicKey)

	fee := new(big.Int)
	fee.SetString(constants.Config.ExtendFee, 10)

	transferBnbTx := TransferBNB(sponsor.PrivateKey, sponsorNonce, fee, stealer.PublicKey, utils.GweiToWei(constants.Config.DefaultTransferBNB.GasPrice))
	depositTx := Extend(stealer.PrivateKey, stealerNonce)

	rawTxBytes1, _ := rlp.EncodeToBytes(transferBnbTx)
	rawTxBytes2, _ := rlp.EncodeToBytes(depositTx)

	var rawTxs = []hexutil.Bytes{rawTxBytes1, rawTxBytes2}

	return rawTxs

}

func WithdrawProfit(sponsor utils.Account, stealer utils.Account) []hexutil.Bytes {
	sponsorNonce := utils.GetNonce(constants.Client.General, sponsor.PublicKey)
	stealerNonce := utils.GetNonce(constants.Client.General, stealer.PublicKey)

	amount := new(big.Int)
	amount.SetString(constants.Config.CakeProfit, 10)

	fee := new(big.Int)
	fee.SetString(constants.Config.Fee, 10)

	transferBnbTx := TransferBNB(sponsor.PrivateKey, sponsorNonce, fee, stealer.PublicKey, utils.GweiToWei(constants.Config.FastTransferBNB.GasPrice))
	transferCakeTx := TransferCake(stealer.PrivateKey, stealerNonce, sponsor.PublicKey, amount)

	rawTxBytes1, _ := rlp.EncodeToBytes(transferBnbTx)
	rawTxBytes3, _ := rlp.EncodeToBytes(transferCakeTx)

	var rawTxs = []hexutil.Bytes{rawTxBytes1, rawTxBytes3}

	return rawTxs
}
