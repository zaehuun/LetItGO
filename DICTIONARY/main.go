package main

import (
	"fmt"

	"github.com/zaehuun/LetItGO/DICTIONARY/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
