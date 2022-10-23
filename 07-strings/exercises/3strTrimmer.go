package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func strTrimmer() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}

func rightTrimmer() {
	name := "inanç           "

	name = strings.TrimRight(name, " ")
	l := utf8.RuneCountInString(name)

	fmt.Println(l)
}
