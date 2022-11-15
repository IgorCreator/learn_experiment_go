package main

import (
	"fmt"
	"strings"
)

func main() {
	iterationWithFullForLoop()

	shortendLooping()
}

func iterationWithFullForLoop() {
	sentence := strings.Fields("a b c d e f g h j k l m n")
	for i := 1; i < len(sentence); i++ {
		fmt.Printf("#%-2d: %s\n", i+1, sentence[i])
	}
}

func shortendLooping() {
	sentence := strings.Fields("z w y u p m n k g q t")
	for i, el := range sentence {
		fmt.Printf("}%-2d: %s\n", i+1, el)
	}

	for i, el := range sentence[5:] {
		fmt.Printf(">%-2d: %s\n", i+1, el)
	}
}
