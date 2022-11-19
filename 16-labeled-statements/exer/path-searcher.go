package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//path := filepath.SplitList(os.Getenv("PATH"))
	path := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	fmt.Println(path)

	query := os.Args[1:]

	for _, q := range query {

	search:
		for i, w := range path {
			q, w = strings.ToLower(q), strings.ToLower(w)
			switch q {
			case "and", "or", "the":
				break search
			}

			if strings.Contains(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue
			}
		}
	}
}
