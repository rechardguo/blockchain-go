package blk

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
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

func (b *Block) SetHash() {
	heightBytes := IntToHex(b.Height)
	fmt.Println(heightBytes)

	timestampStr := strconv.FormatInt(b.Timestamp, 2)
	timestampBytes := []byte(timestampStr)

	blockBytes := bytes.Join([][]byte{heightBytes, b.PrevHash, timestampBytes, b.Data}, []byte{})

	blocksha := sha256.Sum256(blockBytes)

	b.Hash = blocksha[:]
}

func NewBlock(height int64, prevHash []byte, data []byte) *Block {
	block := &Block{height, nil, prevHash, time.Now().Unix(), 0, data}
	block.SetHash()
	return block
}
