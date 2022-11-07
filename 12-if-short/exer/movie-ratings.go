package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run movie-ratings.go 43

func main() {
	mySolutionRating()
}

func mySolutionRating() {
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
	} else if age, err := strconv.Atoi(os.Args[1]); err != nil || age < 0 {
		fmt.Printf("Wrong age: %q\n", os.Args[1])
	} else if age < 13 {
		fmt.Println("PG-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("R-Rated")
	}
}
