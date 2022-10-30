package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

func main() {
	eth_url := os.Getenv("MAINNET_URL")
	client, err := ethclient.Dial(eth_url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
}