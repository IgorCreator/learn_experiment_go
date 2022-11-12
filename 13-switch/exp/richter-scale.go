package main

import (
	"fmt"
	"os"
	"strconv"
)

//  go run richter-scale.go ABC
//  go run richter-scale.go 2
//  go run richter-scale.go 11.0

// EXERCISE: Richter Scale
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
	} else if magnitude, err := strconv.ParseFloat(os.Args[1], 64); err != nil || magnitude < 0.0 {
		fmt.Printf("I couldn't get that, sorry. Input: %q\n", os.Args[1])
	} else {
		richterScale(magnitude)
	}
}

func richterScale(magnitude float64) {
	var desc string

	switch r := magnitude; { // same as -> r := magnitude; switch true {}
	case r > 0.0 && r < 2.0:
		desc = "micro"
	case r > 2.0 && r < 3.0:
		desc = "very minor"
	case r > 3.0 && r < 4.0:
		desc = "minor"
	case r > 4.0 && r < 5.0:
		desc = "light"
	case r > 5.0 && r < 6.0:
		desc = "moderate"
	case r > 6.0 && r < 7.0:
		desc = "strong"
	case r > 7.0 && r < 8.0:
		desc = "major"
	case r > 8.0 && r < 10.0:
		desc = "great"
	default:
		desc = "massive"
	}

	fmt.Printf("%.1f is %s\n", magnitude, desc)
}
