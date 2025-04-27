package main

import (
	blk "blockchain-go/blk" // Import the blk package
	"fmt"
)

func main() {
	block := blk.CreateBlockChainWithGenesisBlock()
	fmt.Println(block.Blocks[0])

}
