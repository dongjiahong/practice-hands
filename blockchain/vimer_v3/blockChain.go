package main

import (
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lashHashKey = "key"

type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB // 数据库操作句柄
	tail []byte   // 尾巴，表示最后一个区块的哈希值
}

func NewBlockChain() *BlockChain {
	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("NewBlockChain 1: ", err)

	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			// 取出最后区块的哈希值
			lastHash = bucket.Get([]byte(lashHashKey))
		} else { // 没有bucket，要去创建创世块，将数据填写到数据库的bucket中
			genesis := NewGenesisBlock()
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			CheckErr("NewBlockChain 2: ", err)
			err = bucket.Put(genesis.Hash, genesis.Serialize()) // 写入区块
			CheckErr("NewBlockChain 3: ", err)
			err = bucket.Put([]byte(lashHashKey), genesis.Hash) // 写入尾节点
			CheckErr("NewBlockChain 4: ", err)
			lastHash = genesis.Hash
		}
		return nil
	})

	return &BlockChain{db, lastHash}
}

func (bc *BlockChain) AddBlock(data string) {
	var prevBlockHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		prevBlockHash = bucket.Get([]byte(lashHashKey))
		return nil
	})

	block := NewBlock(data, prevBlockHash)

	err := bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		err := bucket.Put(block.Hash, block.Serialize())
		CheckErr("AddBlock1 ", err)
		err = bucket.Put([]byte(lashHashKey), block.Hash)
		CheckErr("AddBlock2 ", err)
		bc.tail = block.Hash
		return nil
	})
	CheckErr("AddBlock 3 ", err)
}

type BlockChainIterator struct {
	currHash []byte
	db       *bolt.DB
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{currHash: bc.tail, db: bc.db}
}

func (it *BlockChainIterator) Next() (block *Block) {
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			return nil
		}

		data := bucket.Get(it.currHash)
		block = Deserialize(data)
		it.currHash = block.PrevBlockHash
		return nil
	})
	CheckErr("Next ", err)
	return
}
