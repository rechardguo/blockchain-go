package main

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

func main() {
	// blockchain := blk.CreateBlockChainWithGenesisBlock()

	// blockchain.AddBlockToBlockChain("block1")
	// blockchain.AddBlockToBlockChain("block2")
	// blockchain.AddBlockToBlockChain("block3")
	// blockchain.AddBlockToBlockChain("block4")

	// for _, block := range blockchain.Blocks {
	// 	fmt.Printf("Height: %d\n", block.Height)
	// 	fmt.Printf("Timestamp: %d\n", block.Timestamp)
	// 	fmt.Printf("Nonce: %d\n", block.Nonce)
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Println()
	// }

	// block := blk.NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte("Genesis block"))
	// fmt.Printf("%x\n", block.Hash)
	// fmt.Printf("%d\n", block.Nonce)

	// fmt.Println("--------------------------------")
	// serializedData := block.Serialize()
	// block2 := blk.DeserializeBlock(serializedData)
	// fmt.Printf("%x\n", block2.Hash)
	// fmt.Printf("%d\n", block2.Nonce)

	// Open the database
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a writable transaction
	err = db.Update(func(tx *bbolt.Tx) error {
		var b *bbolt.Bucket
		// Get a bucket
		b = tx.Bucket([]byte("MyBucket"))
		if b == nil {
			// Create a bucket
			bucket, err := tx.CreateBucket([]byte("MyBucket"))
			if err != nil {
				return err
			}
			b = bucket
		}

		// Put a key-value pair in the bucket
		return b.Put([]byte("key"), []byte("value"))
	})

	if err != nil {
		log.Fatal(err)
	}

	// Read the value back
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		v := b.Get([]byte("key"))
		fmt.Printf("Value: %s\n", v)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Delete the key-value pair
	err = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		return b.Delete([]byte("key"))
	})
}
