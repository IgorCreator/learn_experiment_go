package main

import (
	"fmt"
	"strings"
)

func main() {
	const word = "console"

	runiesArr := []byte(word)
	fmt.Printf("%-10s %-10s %-10s %-10s %-12s\n%s\n", "literal", "binary", "dec", "hex", "encoded(utf-8)", strings.Repeat("-", 45))
	for _, ch := range runiesArr {
		fmt.Printf("%-10c 0b%-10[1]b %-10[1]d 0x%-10[1]x % -12x\n", ch, string(ch))
	}

	fmt.Printf("%c\n", runiesArr)
	fmt.Printf("%b\n", runiesArr)
	fmt.Printf("%d\n", runiesArr)
	fmt.Printf("[% x]\n", runiesArr)

	// print the word manually using runes
	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
