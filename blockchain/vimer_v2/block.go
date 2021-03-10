package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Version       int64  // 版本
	PrevBlockHash []byte // 前区块的哈希值
	Hash          []byte // 当前区块的哈希值，为了简化代码
	MerKelRoot    []byte // 梅克尔根
	TimeStamp     int64  // 时间戳
	Bits          int64  // 难度值
	Nonce         int64  // 随机值
	Data          []byte // 交易信息
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		MerKelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          1,
		Nonce:         1,
		Data:          []byte(data),
	}
	block.SetHash()
	return &block
}

func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data,
	}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
