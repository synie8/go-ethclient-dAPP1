package task1

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// QueryByBlockNumber 查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量
func QueryByBlockNumber(client *ethclient.Client, blockNumber *big.Int) (string, uint64, int, error) {
	blockHash := ""
	blockTimestamp := uint64(0)
	blockTranscations := 0

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
		return blockHash, blockTimestamp, blockTranscations, err
	}
	blockHash = block.Hash().Hex()
	blockTimestamp = block.Time()
	blockTranscations = len(block.Transactions())

	fmt.Println("Block Hash: ", blockHash)
	fmt.Println("Block Timestamp: ", blockTimestamp)
	fmt.Println("Block Transcations: ", blockTranscations)

	return blockHash, blockTimestamp, blockTranscations, nil
}
