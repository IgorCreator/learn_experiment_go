package main

import (
	"fmt"
	"os"
	"strings"
)

const sentence = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(sentence)
	query := os.Args[1:]

	fmt.Println("-------------------------------")
	reqExitAsFoundFirstMAtchFromInput(words, query)

	fmt.Println("-------------------------------")
	reqExitAfterFindingOnesInput(words, query)

	fmt.Println("-------------------------------")
	reqSkipFilteredWords(words, query)
}

func reqExitAsFoundFirstMAtchFromInput(words []string, query []string) {
queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				break queries
			}
		}
	}
}

func reqExitAfterFindingOnesInput(words []string, query []string) {
queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// skip duplicate words
				continue queries
			}
		}
	}
}

func reqSkipFilteredWords(words []string, query []string) {
queries:
	for _, q := range query {

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				continue queries
			}
		}
	}
}
