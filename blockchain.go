package main

import (
	"fmt"
	"log"
	"time"
)

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("Chain %d \n", i)
		block.Print()
	}
}

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	// Transforma o timestamp em um valor
	// numérico
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp		%d\n", b.timestamp)
	fmt.Printf("nonce			%d\n", b.nonce)
	fmt.Printf("previousHash		%s\n", b.previousHash)
	fmt.Printf("transactions		%s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	bc := NewBlockChain()
	bc.CreateBlock(2, "init hash")
	bc.Print()
}
