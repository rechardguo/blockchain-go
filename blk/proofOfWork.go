package blk

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const (
	targetBits = 16
)

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	//算出target=0000 0000 0000 0000 1000...0
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) hash(nonce int64) []byte {
	blockBytes := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(pow.Block.Height),
			IntToHex(nonce),
		},
		[]byte{})
	blocksha := sha256.Sum256(blockBytes)
	return blocksha[:]
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hash []byte
	hashInt := big.NewInt(0)
	nonce := int64(0)
	for {
		hash = pow.hash(nonce)
		hashInt.SetBytes(hash)
		fmt.Printf("\r%x\n", hashInt)
		// pow.target=0000 0000 0000 0000 10..0
		// hashInt 的范围可以是下面的范围
		// hashInt   =0000 0000 0000 0000 01..1
		// hashInt   =0000 0000 0000 0000 00..0
		// 上面的算法就是穷举所有的可能，知道得到一个hashInt小于pow.target
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}
	return nonce, hash
}

func (pow *ProofOfWork) Validate() bool {
	hashInt := big.NewInt(0)
	hash := pow.hash(pow.Block.Nonce)
	hashInt.SetBytes(hash)
	return hashInt.Cmp(pow.target) == -1
}
