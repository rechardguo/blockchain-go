package blk

import (
	"log"
	"math/big"

	"go.etcd.io/bbolt"
)

const DB_NAME = "blockchain.db"
const DB_TABLE = "blocks"

type Blockchain struct {
	Tip []byte //最新区块的hash
	DB  *bbolt.DB
}

func (bc *Blockchain) GetBucket(tx *bbolt.Tx) (*bbolt.Bucket, error) {
	var b *bbolt.Bucket
	var err error
	b = tx.Bucket([]byte(DB_TABLE))
	if b == nil {
		b, err = tx.CreateBucket([]byte(DB_TABLE))
		if err != nil {
			return b, err
		}
	}
	return b, err
}

func (bc *Blockchain) AddBlockToBlockChain(data string) {
	var lastestBlk *Block
	bc.DB.View(func(tx *bbolt.Tx) error {
		b, err := bc.GetBucket(tx)
		if err != nil {
			log.Panic(err)
		}
		latestBlkData := b.Get([]byte(bc.Tip))

		lastestBlk = DeserializeBlock(latestBlkData)
		if lastestBlk == nil {
			log.Panicf("latest block with hash=%b is nil", bc.Tip)
		}

		return nil
	})

	newBlock := NewBlock(lastestBlk.Height+1, bc.Tip, []byte(data))
	bc.saveBlockChainToDB(newBlock)
	bc.Tip = newBlock.Hash

}

func (bc *Blockchain) saveBlockChainToDB(block *Block) (bool, error) {
	err := bc.DB.Update(func(tx *bbolt.Tx) error {
		b, err := bc.GetBucket(tx)
		if err != nil {
			log.Panic(err)
		}
		return b.Put(block.Hash, block.Serialize())
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateBlockChainWithGenesisBlock() *Blockchain {
	genesisBlk := CreateGenesisBlock("Genesis Block")
	// Open the database
	db, err := bbolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	bc := Blockchain{genesisBlk.Hash, db}
	bc.saveBlockChainToDB(genesisBlk)
	bc.Tip = genesisBlk.Hash
	return &bc
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bbolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.Tip, bc.DB}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block
	var hashInt big.Int

	hashInt.SetBytes(i.currentHash)

	if hashInt.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	i.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(DB_TABLE))
		blockData := b.Get(i.currentHash)
		block = DeserializeBlock(blockData)
		return nil
	})
	if block != nil {
		i.currentHash = block.PrevHash
	}
	return block
}
