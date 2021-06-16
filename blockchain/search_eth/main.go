package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	token "searcheth/erc"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/3ebcec9d8c6142c9aaa2ba10eec55424")
	if err != nil {
		log.Fatal(err)
	}

	// Golem一个基于以太坊的算力项目代币是GNT, 下面是合约的token地址
	//tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	tokenAddress := common.HexToAddress("0xb8b01cec5ced05c457654fc0fda0948f859883ca")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 这个是账户地址
	//address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	address := common.HexToAddress("0x3e2cf7a1bb4ddb2e4a9513aa8157e43137b3575d")
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("wei: %s\n", balance) // "wei: 74567284359648473625674854"

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
}
