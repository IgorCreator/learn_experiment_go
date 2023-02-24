package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//go run file-reader.go < ./shakespeare.txt

func main() {
	var total int
	words := make(map[string]bool)

	// Print all input
	stdin, _ := os.Open("./shakespeare.txt")
	stdin2, _ := os.Open("./shakespeare.txt")
	in := bufio.NewScanner(stdin)
	in2 := bufio.NewScanner(stdin2)
	for in.Scan() {
		fmt.Printf("%s\n", strings.ToUpper(in.Text()))
	}

	// Call total and unique words
	rx := regexp.MustCompile(`[^A-Za-z]+`)
	in2.Split(bufio.ScanWords)
	for in2.Scan() {
		word := strings.ToLower(in2.Text())
		word = rx.ReplaceAllString(word, "")
		if len(word) > 2 {
			fmt.Printf("%s\n", word)
			words[word] = true
		}
		total++
	}

	fmt.Printf("Total: %d. Unique: %d", total, len(words))
}
