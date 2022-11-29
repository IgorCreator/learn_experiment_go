package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	titles := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) < 2 {
		fmt.Printf("Tell me a book title\n")
		return
	}

	fmt.Println("Search Results:")

	var found bool
	for _, v := range titles {
		if strings.Contains(strings.ToLower(v), strings.ToLower(os.Args[1])) {
			fmt.Println("+", v)
			found = true
		}
	}

	if !found {
		fmt.Printf("We don't have the book: %q\n", os.Args[1])
	}
}
