package task2

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployContractCounter(client *ethclient.Client, gasLimit uint64, privateKeystr string) error {
	//加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeystr)
	if err != nil {
		log.Fatal(err)
	}
	//公钥转换
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//获取地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//获取nonce值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//获取chainId
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//参数封装
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("address => ", address.Hex())
	fmt.Println("tx hash => ", tx.Hash().Hex())

	_ = instance

	return nil
}

// 调用合约方法
func CallContractCounter(client *ethclient.Client, contractAddress string, privateKeystr string) error {

	//合约地址
	contractAddressEth := common.HexToAddress(contractAddress)
	//创建合约实例
	counterContract, err := NewCounter(contractAddressEth, client)
	if err != nil {
		log.Fatal(err)
	}
	//加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeystr)
	if err != nil {
		log.Fatal(err)
	}

	//获取chainId
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 加载私钥，用于发送交易
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(300000)

	// 调用需要交易的方法 (transact)
	tx, err := counterContract.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction confirmed in block: %d\n", receipt.BlockNumber)

	return nil
}
