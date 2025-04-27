package main

import (
	blk "blockchain-go/blk" // Import the blk package
	"fmt"
)

func main() {
	block := blk.CreateGenesisBlock("Genesis Block")
	fmt.Println(block)

}
