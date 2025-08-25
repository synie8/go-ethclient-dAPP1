package main

import (
	"dapp1/task1"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// 链上合约地址
	SEPOLIA_URL = "https://eth-sepolia.g.alchemy.com/v2/z5157D8Ab1rxPnghJcJ1y"
)

func main() {
	client, err := ethclient.Dial(SEPOLIA_URL)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(9058241)
	task1.QueryByBlockNumber(client, blockNumber)

}
