package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("transactions: %s\n", b.transactions)
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 24), i, strings.Repeat("=", 24))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 24))
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}
