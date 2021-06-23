// Package eth 以太坊包
package bzz

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"searcheth/bzz/token"
)

type Config struct {
	NodeApi     string `toml:"NodeApi"`     // eth节点ip
	TokenAddr   string `toml:"TokenAddr"`   // bzz代币合约地址
	AccountAddr string `toml:"AccountAddr"` // 查询账户地址
}

type Client struct {
	conf     *Config
	client   *ethclient.Client
	instance *token.Token
	decimals uint8
	Symbol   string
}

type Event struct {
	From         string  // 发送的地址
	To           string  // 接收的地址
	TxHash       string  // 交易hash
	FromBlockNum uint64  // 查询起始的高度
	ToBlockNum   uint64  // 查询截止的高度
	BlockNum     uint64  // 交易高度
	Value        float64 // 发送的代币
}

func NewClient(conf *Config) (*Client, error) {
	if conf == nil {
		return nil, fmt.Errorf("need config")
	}
	client, err := ethclient.Dial(conf.NodeApi)
	if err != nil {
		return nil, err
	}
	tokenAddress := common.HexToAddress(conf.TokenAddr)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		return nil, err
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &Client{
		conf:     conf,
		client:   client,
		instance: instance,
		decimals: decimals,
		Symbol:   symbol,
	}, nil
}

// GetTokenBalance 获取账户余额
func (c *Client) GetTokenBalance() (*big.Float, error) {
	accountAddress := common.HexToAddress(c.conf.AccountAddr)
	bigBalance, err := c.instance.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		return big.NewFloat(0), err
	}

	fbalance := new(big.Float)
	fbalance.SetString(bigBalance.String())

	balance := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(c.decimals))))

	return balance, nil
}

// GetRecentBlockNum 获取节点最新高度
func (c *Client) GetRecentBlockNum() (uint64, error) {
	header, err := c.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return uint64(0), err
	}
	return header.Number.Uint64(), nil
}

// GetTransferEvent 获取指定高度间的交易记录
func (c *Client) GetTransferEvent(fromBlockNum, toBlockNum uint64) ([]*Event, error) {
	accountAddr := common.HexToAddress(c.conf.AccountAddr)
	contractAddr := common.HexToAddress(c.conf.TokenAddr)
	fmt.Println("------> contractAddr: ", contractAddr, " accountAddr: ", accountAddr, " from: ", fromBlockNum, " to: ", toBlockNum)

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlockNum)),
		ToBlock:   big.NewInt(int64(toBlockNum)),
		Addresses: []common.Address{
			contractAddr,
		},
		Topics: [][]common.Hash{
			{logTransferSigHash},
		},
	}

	logs, err := c.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		fmt.Println("====>>>> no logs <<<<======")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		return nil, err
	}

	var events []*Event

	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			values, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				return nil, err
			}
			if vLog.Removed {
				continue
			}
			if common.HexToAddress(vLog.Topics[2].Hex()) == accountAddr {
				fmt.Println("------> in : ", accountAddr, " tx: ", vLog.TxHash.Hex())
			}
			if common.HexToAddress(vLog.Topics[1].Hex()) == accountAddr {
				fmt.Println("------> out : ", accountAddr, " tx: ", vLog.TxHash.Hex())
			}
			e := &Event{
				From:         common.HexToAddress(vLog.Topics[1].Hex()).Hex(),
				To:           common.HexToAddress(vLog.Topics[2].Hex()).Hex(),
				BlockNum:     vLog.BlockNumber,
				TxHash:       vLog.TxHash.Hex(),
				FromBlockNum: fromBlockNum,
				ToBlockNum:   toBlockNum,
			}
			value := fmt.Sprintf("%s", values[0])
			fValue := new(big.Float)
			fValue.SetString(value)

			e.Value, _ = new(big.Float).Quo(fValue, big.NewFloat(math.Pow10(int(c.decimals)))).Float64()
			events = append(events, e)
		default:
			fmt.Println("unknown sighash: ", vLog.Topics[0].Hex())
		}
	}
	return events, nil
}
