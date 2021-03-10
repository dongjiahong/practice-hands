package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int // 目标值 0x0000010000..00000000
}

const targetBits = 24 // 挖矿难度

func NewProofOfWork(block *Block) *ProofOfWork {

	// 二进制表示 00000000000000000....0001
	target := big.NewInt(1)

	//0x10000000000000000000000000000000000000000000000000000000000
	// 十六进制表示：0x0000010000..00000000
	// int64占64位，每一位用4个字节表示，一共有64*4=256个字节
	target.Lsh(target, uint(256-targetBits)) // 左移256-24=232位,也就是6位，前面5个零
	pow := ProofOfWork{block: block, target: target}
	return &pow
}

func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	block := pow.block
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(nonce),
		block.Data,
	}
	data := bytes.Join(tmp, []byte{})
	return data
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	/* 思路
	1.拼装数据
	2.哈希值转换成big.Int类型
	for nonce {
		hash := sha256(block数据 + nonce)
		if 转换(hash) < pow.target {
			找到了
		} else {
			nonce
		}
	}
	return nonce, hash[:]
	*/

	var hash [32]byte
	var nonce int64 = 0
	var hashInt big.Int

	fmt.Println("Begin Mining...")
	fmt.Printf("target hash: %x\n", pow.target.Bytes())
	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 { // -1表示cmp的小于,具体进去看函数说明
			fmt.Printf("found hash: %x, nonce: %d\n", hash, nonce)
			break
		} else {
			if rand.Intn(1000000) == 999 {
				fmt.Println("now nonce is: ", nonce)
			}
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
