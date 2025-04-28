package blk

import (
	"math/big"

	"go.etcd.io/bbolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bbolt.DB
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
