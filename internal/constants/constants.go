package constants

import (
	"github.com/ethereum/go-ethereum/common"
	"suck_steal/config"
	"suck_steal/internal/utils"
)

var Client = utils.NewClient()
var ContractAbi = utils.ReadAbi("abi.json")
var CakeTransferAbi = utils.ReadAbi("cake.json")
var ContractAddress = common.HexToAddress("0x45c54210128a065de780C4B0Df3d16664f7f859e")
var CakeTransferAddress = common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82")
var Config = config.ReadConfigFile("config.json")
