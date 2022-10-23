package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func repeater() {
	msg := "Hello"
	l := utf8.RuneCountInString(msg)
	s := msg + strings.Repeat("!", l)
	fmt.Println(s)
}
