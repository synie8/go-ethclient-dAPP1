package main

import (
	"dapp1/config"
	"dapp1/task1"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 初始化配置
	cfg := config.GetConfig()

	client, err := ethclient.Dial(cfg.SepoliaURL)
	if err != nil {
		log.Fatal(err)
	}
	// 查询9058241区块信息
	blockNumber := big.NewInt(9058241)
	task1.QueryByBlockNumber(client, blockNumber)
	// 转账eth
	task1.TransferEth(client, cfg.AccountAddress2, float64(0.1), cfg.GasLimit, cfg.PrivateKey)

}
