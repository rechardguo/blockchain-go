package blk

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlockToBlockChain(data string) {
	newBlock := NewBlock(int64(len(bc.Blocks)), bc.Blocks[len(bc.Blocks)-1].Hash, []byte(data))
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateBlockChainWithGenesisBlock() *Blockchain {

	genesisBlk := CreateGenesisBlock("Genesis Block")

	return &Blockchain{[]*Block{genesisBlk}}
}
