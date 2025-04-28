package main

import (
	"blockchain-go/BLK"
	"fmt"
	"time"
)

func main() {
	blockchain := BLK.CreateBlockChainWithGenesisBlock()

	blockchain.AddBlockToBlockChain("block1")
	blockchain.AddBlockToBlockChain("block2")
	blockchain.AddBlockToBlockChain("block3")
	blockchain.AddBlockToBlockChain("block4")

	iterator := blockchain.Iterator()
	for {
		block := iterator.Next()
		if block == nil {
			break
		}
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format(("2006-01-01  15:04:05")))
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
	defer blockchain.DB.Close()
}
