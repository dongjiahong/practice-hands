package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lashHashKey = "key"
const genesisInfo = "genesis  info"

type BlockChain struct {
	//blocks []*Block
	db   *bolt.DB // 数据库操作句柄
	tail []byte   // 尾巴，表示最后一个区块的哈希值
}

func isDBExit() bool {
	_, err := os.Stat(dbFile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func InitBlockChain(address string) *BlockChain {
	if isDBExit() {
		fmt.Println("blockchain exist already, no need to create!")
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("InitBlockChain1: ", err)

	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		coinbase := NewCoinbaseTx(address, genesisInfo)
		genesis := NewGenesisBlock(coinbase)
		bucket, err := tx.CreateBucket([]byte(blockBucket))
		CheckErr("InitBlockChain2: ", err)
		err = bucket.Put(genesis.Hash, genesis.Serialize()) // 写入区块
		CheckErr("InitBlockChain3: ", err)
		err = bucket.Put([]byte(lashHashKey), genesis.Hash) // 写入尾节点
		CheckErr("InitBlockChain4: ", err)
		lastHash = genesis.Hash
		return nil
	})

	return &BlockChain{db, lastHash}
}

func GetBlockChainHandler() *BlockChain {
	if !isDBExit() {
		fmt.Println("create blockchain first!")
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("GetBlockChainHandler1: ", err)

	var lastHash []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			// 取出最后区块的哈希值
			lastHash = bucket.Get([]byte(lashHashKey))
		}
		return nil
	})

	return &BlockChain{db, lastHash}
}

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	var prevBlockHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		prevBlockHash = bucket.Get([]byte(lashHashKey))
		return nil
	})

	block := NewBlock(txs, prevBlockHash)

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

func (bc *BlockChain) FindUTXOTransactions(address string) []Transaction {
	var UTXOTransactions []Transaction

	// 存储使用过的utxo的集合 map[交易id]=>[]int64
	// 0x111111 : 0, 1 都是给Alice转账
	spentUTXO := make(map[string][]int64)

	it := bc.NewIterator()
	for {
		// 便利区块
		block := it.Next()

		// 便利交易
		for _, tx := range block.Transactions {
			//遍历input
			// 目的找到已经消耗过的utxo, 把他们放到一个集合里
			// 需要两个字段来标识使用过的utxo：a.交易id, b.output的索引
			if !tx.IsCoinbase() {
				for _, input := range tx.TXInputs {
					if input.CanUnlockUTXOWith(address) {
						// map[txid][]int64
						spentUTXO[string(tx.TXID)] = append(spentUTXO[string(tx.TXID)], input.Vout)
					}
				}
			}

		OUTPUTS:
			// 遍历output
			// 目的:找到所有能支配的utxo
			for currIndex, output := range tx.TXOutputs {
				// 检查当前output是否已经被消耗，如果消耗过，那么就进行下一个output检验
				if spentUTXO[string(tx.TXID)] != nil {
					// 非空，代表当前交易里面有消耗的utxo
					indexes := spentUTXO[string(tx.TXID)]
					for _, index := range indexes {
						// 当前索引和消耗的索引比较，若相同，表明这个output肯定被消耗了
						if int64(currIndex) == index {
							continue OUTPUTS
						}
					}
				}

				// 如果当前地址是这个utxo的所有者，就满足条件
				if output.CanBeUnlockedWith(address) {
					UTXOTransactions = append(UTXOTransactions, *tx)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return UTXOTransactions
}

// 寻找指定地址能够使用的utxo
func (bc *BlockChain) FindUTXO(address string) []*TXOutput {
	var UTXOs []*TXOutput
	txs := bc.FindUTXOTransactions(address)

	// 遍历交易
	for _, tx := range txs {
		// 遍历output
		for _, utxo := range tx.TXOutputs {
			// 当前地址拥有的utxo
			if utxo.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, &utxo)
			}
		}
	}
	return UTXOs
}
