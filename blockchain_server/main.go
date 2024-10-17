package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockChainServer(uint16(*port))
	app.Run()

	// imprime o endereço de memória
	// da porta
	// fmt.Println(port)
	// fmt.Println(*port) imprime o número da porta

}
