package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}
	fmt.Println("Original:", items)

	// sorting the mid will affect all items
	mid := len(items) / 2
	midEl := items[mid-1 : mid+2]
	sort.Strings(midEl)
	fmt.Println("Mid:", midEl)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
