package main

import (
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int // 目标值 0x00000010000..00000000
}

const targetBits = 24

func NewProofOfWork(block *Block) *ProofOfWork {

	// 00000000000000000....0001
	target := big.NewInt(1)

	// 0x00000010000..00000000
	target.Lsh(target, uint(256-targetBits)) // 左移256-24=232位
	pow := ProofOfWork{block: block, target: target}
	return &pow
}
