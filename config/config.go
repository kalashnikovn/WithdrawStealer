package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Sponsor            string `json:"sponsor"`
	Stealer            string `json:"stealer"`
	PublicRPC          string `json:"publicRPC"`
	CakeAmount         string `json:"cakeAmount"`
	CakeProfit         string `json:"cakeProfit"`
	PreCallSec         int64  `json:"preCallSec"`
	Fee                string `json:"fee"`
	ExtendFee          string `json:"extendFee"`
	DefaultTransferBNB struct {
		GasPrice int64 `json:"gasPrice"`
		GasLimit int64 `json:"gasLimit"`
	} `json:"defaultTransferBNB"`
	FastTransferBNB struct {
		GasPrice int64 `json:"gasPrice"`
		GasLimit int64 `json:"gasLimit"`
	} `json:"fastTransferBNB"`
	WithdrawAll struct {
		GasPrice int64 `json:"gasPrice"`
		GasLimit int64 `json:"gasLimit"`
	} `json:"withdrawAll"`
	TransferCake struct {
		GasPrice int64 `json:"gasPrice"`
		GasLimit int64 `json:"gasLimit"`
	} `json:"transferCake"`
	Deposit struct {
		ExtendTime int64 `json:"extendTime"`
		GasPrice   int64 `json:"gasPrice"`
		GasLimit   int64 `json:"gasLimit"`
	} `json:"deposit"`
}

func ReadConfigFile(filePath string) Config {
	configFile, err := os.Open(filePath)
	if err != nil {
		return Config{}
	}
	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		return Config{}
	}

	return config
}
