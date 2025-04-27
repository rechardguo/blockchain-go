package main

import (
	blk "blockchain-go/blk" // Import the blk package
	"fmt"
)

func main() {
	block := blk.NewBlock(1, []byte{}, []byte("Genesis Block")) // Use the NewBlock function from the blk package
	fmt.Println(block.Hash)

}
