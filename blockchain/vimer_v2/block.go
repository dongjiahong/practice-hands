package main

import (
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
		Bits:          targetBits,
		Nonce:         0,
		Data:          []byte(data),
	}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
