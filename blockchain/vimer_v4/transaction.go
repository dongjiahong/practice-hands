package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"os"
)

const reward = 12.5

type Transaction struct {
	TXID      []byte     // 交易ID
	TXInputs  []TXInput  // 输入
	TXOutputs []TXOutput // 输出
}

type TXInput struct {
	TXID      []byte // 所引用输出的交易ID
	Vout      int64  // 所引用output的索引值
	ScriptSig string // 解锁脚本，指明可以使用某个output的条件
}

// 检查当前用户能否解开应用的utxo
func (input *TXInput) CanUnlockUTXOWith(unlockData string) bool {
	return input.ScriptSig == unlockData
}

type TXOutput struct {
	Value        float64 // 支付给收款方的金额
	ScriptPubKey string  // 锁定脚本，指定收款方的地址
}

// 检查当前用户是否是这个utxo的所有者
func (output *TXOutput) CanBeUnlockedWith(unlockData string) bool {
	return output.ScriptPubKey == unlockData
}

func (tx *Transaction) SetTXID() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	CheckErr("SetTXID", err)
	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
}

func NewCoinbaseTx(address string, data string) *Transaction {
	if len(data) == 0 {
		data = fmt.Sprintf("reward to %s %f btc", address, reward)
	}

	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}

	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetTXID()
	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	if len(tx.TXInputs) == 1 {
		if len(tx.TXInputs[0].TXID) == 0 && tx.TXInputs[0].Vout == -1 {
			return true
		}
	}
	return false
}
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {

	// map[string][]int64 key:交易id, value:引用output的索引数组
	validUTXOs := make(map[string][]int64)
	var total float64
	validUTXOs /*所需要的，合理的utxo的合集*/, total /*返回utxo的金额总和*/ = bc.FindSuitableUTXOs(from, amount)

	// validUTXOs[0x11111111] = []int64{1}
	// validUTXOs[0x22222222] = []int64{0}
	// ....
	// validUTXOs[0xnnnnnnnn] = []int64{0, 4, 8}
	if total < amount {
		fmt.Println("Not enough money!")
		os.Exit(1)
	}

	var inputs []TXInput
	var outputs []TXOutput

	// 1.创建inputs
	// 进行output到input的转换
	// 遍历有效的utxo的合集
	for txId, outputIndexes := range validUTXOs {
		// 遍历所有引用的utxo的索引，每一个索引需要创建一个input
		for _, index := range outputIndexes {
			input := TXInput{[]byte(txId), int64(index), from}
			inputs = append(inputs, input)
		}
	}

	// 2.创建outputs
	// 给对方支付
	output := TXOutput{amount, to}
	outputs = append(outputs, output)

	// 找零
	if total > amount {
		output := TXOutput{total - amount, from}
		outputs = append(outputs, output)
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetTXID()
	return &tx
}
