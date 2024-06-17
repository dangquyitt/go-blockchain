package main

import (
	"fmt"
	"log"
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
	fmt.Printf("nonce\t\t\t%d\n", b.nonce)
	fmt.Printf("previousHash\t\t\t%s\n", b.previousHash)
	fmt.Printf("timestamp\t\t\t%d\n", b.timestamp)
	fmt.Printf("transactions\t\t\t%s\n", b.transactions)
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
