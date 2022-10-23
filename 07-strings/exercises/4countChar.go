package main

import (
	"fmt"
	"unicode/utf8"
)

func countChar() {
	str := "péripatéticien"
	length := utf8.RuneCountInString(str)
	fmt.Println(len(str))
	fmt.Println(length)
}
