package main

import "fmt"

func main() {
	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}
	var bwords [][]byte

	for _, word := range words {
		wordBytes := []byte(word)
		fmt.Printf("%v\n", wordBytes)
		bwords = append(bwords, wordBytes)
	}

	for _, wordB := range bwords {
		fmt.Printf("%s\n", string(wordB))
	}
}
