package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
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
func NewTransaction(form, to string, amount float64, bc *BlockChain) *Transaction {
	return nil
}
