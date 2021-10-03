package main

import (
	"fmt"
	"log"

	"github.com/zaehuun/LetItGO/DICTIONARY/accounts"
)

func main() {
	account := accounts.NewAccount("zaehuun")
	account.Deposit(20)
	err := account.Withdraw(20)
	fmt.Println(account)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance(), account.Owner())
}
