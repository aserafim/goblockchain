package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	// Transforma o timestamp em um valor
	// numérico
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp		%d\n", b.timestamp)
	fmt.Printf("nonce			%d\n", b.nonce)
	fmt.Printf("previousHash		%x\n", b.previousHash)

	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Timestamp    int64          `json:"timestamp"`
			Nonce        int            `json:"nonce"`
			PreviousHash [32]byte       `json:"previous_hash"`
			Transactions []*Transaction `json:"transactions"`
		}{
			Timestamp:    b.timestamp,
			Nonce:        b.nonce,
			PreviousHash: b.previousHash,
			Transactions: b.transactions,
		})
}

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockChainAddress string
}

func NewBlockChain(blockChainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockChainAddress = blockChainAddress
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
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockChainAddress,
				t.recipientBlockChainAddress,
				t.value))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockChainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) CalculateTotalAmount(blockChainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if blockChainAddress == t.recipientBlockChainAddress {
				totalAmount += value
			}

			if blockChainAddress == t.senderBlockChainAddress {
				totalAmount -= value
			}
		}
	}

	return totalAmount
}

type Transaction struct {
	senderBlockChainAddress    string
	recipientBlockChainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s Transaction %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	//fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address		%s\n", t.senderBlockChainAddress)
	fmt.Printf(" sender_blockchain_address		%s\n", t.recipientBlockChainAddress)
	fmt.Printf(" sender_blockchain_address		%.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			SenderBlockChainAddress    string  `json:"sender_blockchain_address"`
			RecipientBlockChainAddress string  `json:"sender_blockchain_address"`
			Value                      float32 `json:"sender_blockchain_address"`
		}{
			SenderBlockChainAddress:    t.senderBlockChainAddress,
			RecipientBlockChainAddress: t.recipientBlockChainAddress,
			Value:                      t.value,
		})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "my_blockchain_address"
	bc := NewBlockChain(myBlockchainAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	bc.Print()

	bc.AddTransaction("C", "D", 3.0)
	bc.AddTransaction("C", "E", 1.0)
	bc.Mining()
	bc.Print()

	fmt.Printf("C %.1f\n", bc.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", bc.CalculateTotalAmount("D"))
	fmt.Printf("A %.1f\n", bc.CalculateTotalAmount("A"))
	fmt.Printf("B %.1f\n", bc.CalculateTotalAmount("B"))

	/* prevHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, prevHash)

	bc.AddTransaction("C", "D", 3.0)
	bc.AddTransaction("C", "E", 1.0)
	prevHash = bc.LastBlock().Hash()
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce, prevHash)

	bc.AddTransaction("E", "W", 0.5)
	prevHash = bc.LastBlock().Hash()
	nonce = bc.ProofOfWork()
	bc.CreateBlock(nonce, prevHash)
	bc.Print() */

}
