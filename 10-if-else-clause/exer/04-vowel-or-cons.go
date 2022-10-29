package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	firstSol()

	secondSol()
}

func secondSol() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	s := args[1]
	if strings.IndexAny(s, "aeiou") != -1 {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}

func firstSol() {
	input := os.Args[1]

	if len(os.Args) != 2 {
		fmt.Println("Give me only one letter as arg")
	} else {
		if len(input) == 1 {
			i := []rune(input)
			recognizeChar(i[0])
		} else {
			fmt.Println("Give me a letter")
		}
	}
}

func recognizeChar(character rune) {
	if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' {
		fmt.Printf(" %c is vowel.\n", character)
	} else if character == 'y' || character == 'w' {
		fmt.Printf(" %c is sometimes a vowel, sometimes not.\n", character)
	} else {
		fmt.Printf(" %c is consonant.\n", character)
	}
}
