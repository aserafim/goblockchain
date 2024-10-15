package main

import (
	"fmt"
	"log"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	log.Println("teste 1")
	fmt.Println("Teste 2")
}
