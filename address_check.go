package main

import (
	"context"
    "fmt"
    "log"
    "regexp"
	"os"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	eth_url := os.Getenv("MAINNET_URL")
	client, err := ethclient.Dial(eth_url)
	if err != nil {
		log.Fatal(err)
	}

	checkIsContract("0xe41d2489571d322189246dafa5ebde1f4699f498", client)

	checkIsContract("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4", client)
}

func checkIsContract(_address string, client *ethclient.Client) {
	address := common.HexToAddress(_address)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)
}