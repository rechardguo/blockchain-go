package main

import (
	blk "blockchain-go/blk" // Import the blk package
	"fmt"
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

	block := blk.NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte("Genesis block"))
	fmt.Printf("%x\n", block.Hash)
	fmt.Printf("%d\n", block.Nonce)

	fmt.Println("--------------------------------")
	serializedData := block.Serialize()
	block2 := blk.DeserializeBlock(serializedData)
	fmt.Printf("%x\n", block2.Hash)
	fmt.Printf("%d\n", block2.Nonce)
}
