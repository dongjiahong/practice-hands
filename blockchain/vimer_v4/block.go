package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	Version       int64          // 版本
	PrevBlockHash []byte         // 前区块的哈希值
	Hash          []byte         // 当前区块的哈希值，为了简化代码
	MerKelRoot    []byte         // 梅克尔根
	TimeStamp     int64          // 时间戳
	Bits          int64          // 难度值
	Nonce         int64          // 随机值
	Transactions  []*Transaction // 交易信息
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr("Serialize ", err)

	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {
	if len(data) == 0 {
		return nil
	}

	var block Block
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&block)
	CheckErr("Deserialize ", err)
	return &block
}

func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		MerKelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          targetBits,
		Nonce:         0,
		Transactions:  txs,
	}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (block *Block) TransactionHash() []byte {
	var txHashs [][]byte
	txs := block.Transactions
	// 遍历交易
	for _, tx := range txs {
		txHashs = append(txHashs, tx.TXID)
	}

	// 对二维切片进行拼接，生成一维切片
	data := bytes.Join(txHashs, []byte{})
	hash /*[32]byte*/ := sha256.Sum256(data)
	return hash[:]
}
