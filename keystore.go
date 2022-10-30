package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "os"

    "github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "wallets/UTC--2022-10-22T16-35-06.227741641Z--f9c6394c9a3a2eada7cbbc31225355bae56fb674"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// createKs()
	importKs()
}