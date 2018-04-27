package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index         int64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	index := []byte(strconv.FormatInt(b.Index, 10))
	headers := bytes.Join([][]byte{index, timestamp, b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func (b *Block) String() string{
	return fmt.Sprintf("Index: %d\nData: %s\nPrevHash: %x\nHash: %x\n", b.Index, b.Data, b.PrevBlockHash, b.Hash)
}

func NewBlock(data string, index int64, prevBlockHash []byte) *Block {
	block := &Block{index + 1,
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{}}

	block.SetHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(data, prevBlock.Index, prevBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", 0, []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{blocks: []*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")

	for _, b := range bc.blocks {
		fmt.Println(b)
	}
}
