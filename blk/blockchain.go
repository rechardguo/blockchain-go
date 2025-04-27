package blk

type Blockchain struct {
	Blocks []*Block
}

func CreateBlockChainWithGenesisBlock() *Blockchain {

	genesisBlk := CreateGenesisBlock("Genesis Block")

	return &Blockchain{[]*Block{genesisBlk}}
}
