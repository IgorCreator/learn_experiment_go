package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"strconv"
)

// ---------------------------------------------------------
// Observe the length and capacity
// ---------------------------------------------------------

func main() {
	s.PrintBacking = true
	// --- #1 ---
	// 1. create a new slice named: games
	var games []string
	// 2. print the length and capacity of the games slice
	s.Show("Cap and len", games)

	// 3. declare it as an empty slice
	gamesExp1 := []string(nil)
	// 4. print the length and capacity of the games slice
	s.Show("Cap and len", gamesExp1)
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	gamesExp1 = append(gamesExp1, "pacman", "mario", "tetris", "doom")
	// 6. print the length and capacity of the games slice
	s.Show("Cap and len", gamesExp1)

	// 8. declare the games slice again using a slice literal, loop and print iterations
	var gamesExp2 []string
	for i, el := range gamesExp1 {
		gamesExp2 = append(gamesExp2, el)
		s.Show("Cap and len", gamesExp2)
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i+1, len(gamesExp2), cap(gamesExp2))
	}
	fmt.Println()

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := gamesExp1[:0]
	// 2. print the games and the new slice's len and cap
	s.Show("Cap and len", gamesExp1)
	s.Show("Cap and len", zero)
	// 3. append a new element to the new slice
	newSlice := append(zero, "new elem")
	// 4. print the new slice's lens and caps
	s.Show("Cap and len", newSlice)
	// 5. repeat the last two steps 5 times (use a loop)
	for i := 0; i < 5; i++ {
		charIdx := strconv.Itoa(i)
		newSlice = append(newSlice, "new elem idx = "+charIdx)
		s.Show("Cap and len", newSlice)
	}
	// 6. notice the growth of the capacity after the 5th append

	// Use this slice's elements to append to the new slice:
	newSlice = append(newSlice, []string{"ultima", "dagger", "pong", "coldspot", "zetra"}...)
	s.Show("Cap and len", newSlice)
	fmt.Println()

	// zero := ...
	fmt.Printf("games's     len: %d cap: %d\n", len(gamesExp1), cap(gamesExp1))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	s.Show("Cap and len gamesExp1:", gamesExp1)
	s.Show("Cap and len", zero)

	// --- #4 ---
	fmt.Println()

	for n := range zero {
		str := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(str), cap(str))
	}

	// --- #5 ---
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		str := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(str), cap(str), str)
	}

	// --- #6 ---
	fmt.Println()
	zero[0] = "command & conquer"
	gamesExp1[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", gamesExp1)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	// uncomment and see the error.
	// _ = games[:cap(games)+1]
	// or:
	// _ = games[:5]
}
