package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// go run main.go "$(cat ./text.txt)"

var (
	lineCharLimit = 40
)

func main() {
	data := os.Args[1]
	//fmt.Printf("%s %d \n", data, len(data))
	if len(data) < 1 {
		fmt.Println("Please provide text as input")
		fmt.Println("To run: [go run main.go \"$(cat ./text.txt)\"]")
		fmt.Println("To run: [go run main.go \"A lot of text right here to process and wrap\"]")
		return
	}

	splitBySpaces(data)

	fmt.Println()
	processByRunes(data)
}

func processByRunes(data string) {
	var lw int // line width

	for _, r := range data {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > lineCharLimit && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
	fmt.Println()
}

func splitBySpaces(data string) {
	words := strings.Fields(data)
	var buffer bytes.Buffer
	var memory int
	for _, word := range words {
		size := utf8.RuneCountInString(word)
		memory += size

		if memory > lineCharLimit {
			memory = 0
			buffer.WriteString("\n")
		}

		buffer.WriteString(word + " ")
		//fmt.Printf("%s - sym: %d, runes: %d.\n", word, len(word), size)
	}

	fmt.Println(buffer.String())
}
