package main

import (
	blk "blockchain-go/blk" // Import the blk package
	"fmt"
)

func main() {
	blockchain := blk.CreateBlockChainWithGenesisBlock()
	fmt.Println(blockchain.Blocks[0])

	blockchain.AddBlockToBlockChain("block1")
	blockchain.AddBlockToBlockChain("block2")
	blockchain.AddBlockToBlockChain("block3")
	blockchain.AddBlockToBlockChain("block4")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
