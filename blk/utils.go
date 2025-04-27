package blk

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte {

	b := make([]byte, 8)
	for i := 0; i < 8; i++ {
		b[i] = byte(num >> uint(56-8*i) & 0xff)
	}
	return b
}

func IntToHex2(num int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
