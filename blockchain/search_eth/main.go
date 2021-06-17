package main

import (
	"context"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	token "searcheth/erc"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://mainnet.infura.io/v3/3ebcec9d8c6142c9aaa2ba10eec55424")
	if err != nil {
		panic(err)
	}
}

// 获取代币余额
func getTokenBalance(tokenAddr, accountAddr string) (*big.Float, error) {
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		return big.NewFloat(0), err
	}

	address := common.HexToAddress(accountAddr)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return big.NewFloat(0), err
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("name: %s, symbol: %s, decimals: %d\n", name, symbol, decimals)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(decimals))))
	return value, nil
}

// 获取节点的最新高度
func getBlockHeight() (*big.Int, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return header.Number, nil
}

func getTransferEvent(tokenAddr string) {
	// 获取最新高度
	height, err := getBlockHeight()
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(tokenAddr)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(12649414), // 查询的其实高度
		ToBlock:   height,               // 查询截止高度
		Addresses: []common.Address{
			contractAddr, // 要查询的合约
		},
	}

	// 查询事件日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 序列号abi
	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	// 获取查询方法的签名，去结合etherscan.io查看每个代币的合约
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	log.Println("----> sig: ", logTransferSigHash)

	for _, vLog := range logs {
		log.Println("Log block number: ", vLog.BlockNumber, " Index: ", vLog.Index)
		// transfer函数有三个topic
		// 第一个topics[0]是函数名的签名,
		// 第二个topics[1]是第一个参数即，发送的地址
		// 第三个topics[2]是第二个参数即，接收的地址
		/*
			[topic0] 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
			[topic1] 0x00000000000000000000000033303c7bb968d9066eda76da902bead843ab84ad
			[topic2] 0x00000000000000000000000077696bb39917c91a0c3908d577d5e322095425ca
		*/
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			log.Println("Log Name: Transfer")

			event, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatalln(err)
			}
			// event就是交易的值，即所转的代币
			log.Println("event size: ", len(event), " event[0]: ", event[0])

		default:
			log.Println("info: ", vLog.Topics[0].Hex())
		}
	}
}

func main() {
	// gnt token合约地址: 0xa74476443119A942dE498590Fe1f2454d7D4aC0d
	// gnt 要查询的账户地址: 0x614055249E6B330F34E52de7415439E6919d3A46
	tokenAddr := "0xa74476443119A942dE498590Fe1f2454d7D4aC0d"
	accountAddr := "0x614055249E6B330F34E52de7415439E6919d3A46"
	balance, err := getTokenBalance(tokenAddr, accountAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("balance: %f\n", balance)

	getTransferEvent(tokenAddr)
}
