package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const MINING_DIFFICULTLY = 3

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
		transactions: transactions,
	}
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previous_hash: %x\n", b.previousHash)
	fmt.Printf("timestamp: %d\n", b.timestamp)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 24), i, strings.Repeat("=", 24))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 24))
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, len(bc.transactionPool))
	for i, t := range bc.transactionPool {
		transactions[i] = NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value)
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficutly int) bool {
	zeros := strings.Repeat("0", difficutly)
	guessBlock := Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    0,
		transactions: transactions,
	}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficutly] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTLY) {
		nonce += 1
	}
	return nonce
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		senderBlockchainAddress:    sender,
		recipientBlockchainAddress: recipient,
		value:                      value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 48))
	fmt.Printf("	sender_blockchain_address: %s\n", t.senderBlockchainAddress)
	fmt.Printf("	recipient_blockchain_address: %s\n", t.recipientBlockchainAddress)
	fmt.Printf("	value: %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().Hash()
	nonce := blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	previousHash = blockchain.LastBlock().Hash()
	nonce = blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()
}
