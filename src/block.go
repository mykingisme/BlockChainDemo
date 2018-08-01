package main

import (
	"bytes"
	"crypto/sha256"

	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	TimeStamp     int64
	TargetBits    int64
	Nonce         int64
	merKerRoot    []byte
	Data          []byte
}

func CreateBlock(data string, prevBlockHash []byte) (block *Block) {
	block = &Block{
		1,
		prevBlockHash,
		[]byte{},
		time.Now().Unix(),
		10,
		10,
		[]byte{},
		[]byte(data)}
	block.SetHash()

	return
}
func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.Data,
		IntToByte(block.Nonce),
		block.merKerRoot}
	data := bytes.Join(tmp, []byte{})
	temp := sha256.Sum256(data)
	block.Hash = temp[:]
}

func CreateGenesisBlock() *Block {
	return CreateBlock("The first block", nil)
}
