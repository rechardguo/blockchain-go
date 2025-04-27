package blk

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Height    int64
	Hash      []byte
	PrevHash  []byte
	Timestamp int64
	Nonce     int64
	Data      []byte
}

func NewBlock(height int64, prevHash []byte, data []byte) *Block {
	block := &Block{height, nil, prevHash, time.Now().Unix(), 0, data}
	//调用工作量证明返回有效的Hash和Nonce
	block.Nonce, block.Hash = NewProofOfWork(block).Run()
	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []byte(data))
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
