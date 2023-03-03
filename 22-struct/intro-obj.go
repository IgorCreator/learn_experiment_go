package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn  string
		title string
	}

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	moby.text.words = 1000
	moby.words++
	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, // equals to: moby.text.title
		moby.words, // equals to: moby.text.words
		moby.isbn)
	fmt.Printf("%#v\n", moby)

	spew.Dump(moby)
}
