package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("A send B 1BTC")
	bc.AddBlock("B send C 1BTC")

	for _, block := range bc.blocks {
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Bits: %d\n", block.Bits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("IsValid: %v\n", NewProofOfWork(block).IsValid())
		fmt.Println("----------------------------------------------")
	}
}
