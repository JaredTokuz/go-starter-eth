package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"context"
)

func main() {
	eth_url := os.Getenv("MAINNET_URL")
	client, err := ethclient.Dial(eth_url)
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	blockNumber := big.NewInt(5532993)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}