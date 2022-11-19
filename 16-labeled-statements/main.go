package main

import (
	"fmt"
	"os"
	"strings"
)

const statement = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(statement)
	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-d: %q\n", i+1, q)
				break
			}
		}
	}
}
