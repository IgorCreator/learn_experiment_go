package main

import (
	s "github.com/inancgumus/prettyslice"
)

// Difference between length and capacity in Slice?
//
// The length describes the length of a slice but a
// capacity describes the length of the backing array
// beginning from the first element of the slice

func main() {
	s.PrintBacking = true

	// ----------------------------------------------------
	// #1 nil slice
	var games []string // nil slice
	s.Show("games", games)

	// ----------------------------------------------------
	// #2 empty slice
	games = []string{} // empty slice
	s.Show("games", games)
	s.Show("look at ptr this one and prev", []int{})

	// ----------------------------------------------------
	// #3 non-empty slice
	games = []string{"pacman", "mario", "tetris", "doom"}
	s.Show("games", games)

	// ----------------------------------------------------
	// #4 reset the part using the games slice
	//    part is empty but its cap is still 4
	part := games
	s.Show("part", part)

	part = games[:0]
	s.Show("part[:0]", part)
	s.Show("part[:cap]", part[:cap(part)])

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	part = games
	for cap(part) != 0 {
		part = part[1:len(part)]
		s.Show("part", part)
	}

	// #6 backing array's elements become inaccessible
	// games = games[len(games):]

	// ----------------------------------------------------
	// #5 part doesn't have any more capacity
	//    games slice is still intact
	s.Show("part", part)
	s.Show("games", games)
}
