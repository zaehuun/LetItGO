package main

import (
	"log"

	"github.com/zaehuun/LetItGO/DICTIONARY/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
}
