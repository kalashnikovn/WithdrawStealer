package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"suck_steal/internal/constants"
	"suck_steal/internal/transactions"
	"suck_steal/internal/utils"
	"time"
)

func main() {

	sponsor := utils.NewAccount(constants.Config.Sponsor)
	stealer := utils.NewAccount(constants.Config.Stealer)
	sponsorNonce := utils.GetNonce(constants.Client.General, sponsor.PublicKey)
	timestamp := utils.GetTimestamp(stealer.PublicKey, "https://rpc.ankr.com/bsc")

	fmt.Println("Timestamp: ", timestamp)
	fmt.Printf("Waiting for: %v\n", time.Unix(timestamp, 0))

	WaitAndRun(timestamp, func() {
		Withdraw(sponsor, stealer, sponsorNonce)
	})

}

func executeAfterTimestamp(publicKey common.Address, functionToExecute func()) error {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	rpc := constants.Config.PublicRPC

	timestampCh := make(chan int64)
	go func() {
		for {
			timestamp := utils.GetTimestamp(publicKey, rpc)
			fmt.Println(timestamp)
			timestampCh <- timestamp
		}
	}()

	for {
		select {
		case timestamp := <-timestampCh:
			currentTime := time.Now().Unix()
			if currentTime >= timestamp {
				functionToExecute()
				return nil
			}
		case <-ticker.C:
			// Do nothing, just wait
		}
	}
}

func WaitAndRun(timestamp int64, f func()) {
	now := time.Now().Unix()
	if now > timestamp {
		f()
		return
	}

	duration := time.Duration(timestamp-now) * time.Second
	//halfSecond := time.Duration(500) * time.Millisecond
	fmt.Println("Wait:", duration)
	time.Sleep(duration)
	f()
}

func Withdraw(sponsor, stealer utils.Account, sponsorNonce uint64) error {
	var rawTxs = transactions.WithdrawFullCake(sponsor, stealer, sponsorNonce)
	res, err := transactions.SendBundleTx(rawTxs)
	if err != nil {
		return err
	}

	log.Println(res)
	return nil

}

func WithdrawWithNonce(sponsor, stealer utils.Account, stealerNonce uint64, sponsorNonce uint64) {
	var rawTxsPlusNonce = transactions.WithdrawFullCakePlusNonce(sponsor, stealer, stealerNonce, sponsorNonce)
	res, err := transactions.SendBundleTx(rawTxsPlusNonce)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Plus Nonce: ", res)
}

func Extend(sponsor, stealer utils.Account) {
	var rawTxs = transactions.ExtendTime(sponsor, stealer)
	res, err := transactions.SendBundleTx(rawTxs)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println(res)
}
