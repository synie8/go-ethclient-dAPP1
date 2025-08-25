package task1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

// eth 转账
func TransferEth(client *ethclient.Client, toAddressStr string, amount float64, gasLimit uint64, privateKeystr string) error {
	//加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeystr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return err
	}
	//发送地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//交易随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//eth数量转 wei
	amountBig := new(big.Float).SetFloat64(amount)
	amountWei := new(big.Float).Mul(amountBig, big.NewFloat(1000000000000000000))
	value := new(big.Int)
	amountWei.Int(value)
	//交易gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}
	//接收地址
	toAddress := common.HexToAddress(toAddressStr)
	//构建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	//签名交易
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("交易发送成功")
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	return nil
}
