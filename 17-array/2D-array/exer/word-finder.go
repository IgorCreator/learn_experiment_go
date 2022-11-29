package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Plerunase give me numbers\n")
		return
	}
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		q = strings.ToLower(q)

		for i, w := range words {
			switch q {
			case "and", "or", "the", "was", "since", "very":
				continue
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
