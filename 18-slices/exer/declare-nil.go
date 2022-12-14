package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
)

func main() {
	myNoviceLvl()
	fmt.Printf("----------------\n")
	teacherEx()
}

func teacherEx() {
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
}

func myNoviceLvl() {
	var friends []string
	var distance []int
	var data []byte
	var ratios []float64
	var alives []bool

	s.Show("friends", friends)
	s.Show("distance", distance)
	s.Show("data", data)
	s.Show("ratios", ratios)
	s.Show("alives", alives)
}
