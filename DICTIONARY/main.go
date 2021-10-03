package main

import (
	"fmt"

	"github.com/zaehuun/LetItGO/DICTIONARY/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary.Add("hello", "first")
	dictionary.Search("hello")
	dictionary.Delete("hello")
	word, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	// dictionary.Update()
}
