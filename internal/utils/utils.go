package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/bnb48club/puissant_sdk/bnb48.sdk"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math"
	"math/big"
)

func NewClient() *bnb48.Client {
	client, _ := bnb48.Dial("https://fonce-bsc.bnb48.club", "https://puissant-bsc.bnb48.club")
	return client
}

func ReadAbi(filename string) abi.ABI {
	abiBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to load ABI file: %v", err)
	}

	contractAbi, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	return contractAbi
}

func StrToWallet(privatekey string) (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	return privateKey, fromAddress
}

func GetNonce(client *ethclient.Client, fromAddress common.Address) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	return nonce

}

func ReadContract(address common.Address, rpc string) []interface{} {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress("0x45c54210128a065de780C4B0Df3d16664f7f859e")

	contractAbi := ReadAbi("abi.json")

	input, err := contractAbi.Pack("userInfo", address)
	if err != nil {
		panic(err)
	}

	data, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	}, nil)
	if err != nil {
		panic(err)
	}

	result, err := contractAbi.Unpack("userInfo", data)
	if err != nil {
		panic(err)
	}

	return result
}

func ToInt64(n *big.Int) (int64, error) {
	if n.IsInt64() {
		return n.Int64(), nil
	}

	if n.IsUint64() && n.Uint64() <= uint64(int64(^uint64(0)>>1)) {
		return int64(n.Uint64()), nil
	}

	return 0, fmt.Errorf("невозможно сконвертировать *big.Int в int64: %v", n)
}

func GweiToWei(n int64) int64 {
	return n * 1000000000
}

func ToWei(n float64) int64 {
	return int64(math.Round(n * 1e18))
}

func GetTimestamp(account common.Address, rpc string) int64 {
	result := ReadContract(account, rpc)
	timestamp, _ := ToInt64(result[5].(*big.Int))

	return timestamp
}
