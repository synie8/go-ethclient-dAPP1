package main

import (
	"dapp1/config"
	"dapp1/task2"
	"log"

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
	//blockNumber := big.NewInt(9058241)
	//task1.QueryByBlockNumber(client, blockNumber)
	// 转账eth
	//task1.TransferEth(client, cfg.AccountAddress2, float64(0.1), cfg.GasLimit, cfg.PrivateKey)
	//部署合约
	//fmt.Println("部署合约 counter ************************")
	//task2.DeployContractCounter(client, cfg.GasLimit, cfg.PrivateKey)
	//部署成功的合约地址
	contractAddress := "0x8f656D2CD9EEd6A4Aa543a99CDc890926Ad66f19"
	// 调用合约方法
	task2.CallContractCounter(client, contractAddress, cfg.PrivateKey)

}
