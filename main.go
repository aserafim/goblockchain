package main

import (
	"fmt"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockChainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockChainAddress(), "B", 1.0)
	fmt.Printf("signature %s\n", t.GenerateSignature())

}
