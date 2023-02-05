package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)

	// ~~~ CHANGE THIS CODE ~~~
	//
	fix := make([]string, len(lyric)+3)
	cutpoints := []int{8, 10, 5}
	for i, n := 0, 0; n < len(lyric); i++ {
		n += copy(fix[n+i:], lyric[n:n+cutpoints[i]])
		fix[n+i] = "\n"
	}

	// ===================================
	// WRONG: impl
	//var fix, part1, part2, part3 []string
	//
	//copy(part1, lyric[:8])
	//copy(part2, lyric[8:18])
	//copy(part3, lyric[18:23])
	//
	//part1 = append(part1, "\n")
	//part2 = append(part2, "\n")
	//part3 = append(part3, "\n")
	//s.Show("fix slice", lyric[:8])
	//s.Show("fix slice", part1)
	//s.Show("fix slice", part2)
	//s.Show("fix slice", part3)
	//
	//copy(fix, part1)
	//copy(fix, part2)
	//copy(fix, part3)
	//
	// ===================================

	// Currently, it prints every sentence on the same line.
	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
