package main

import (
	"fmt"
	"os"
)

//  go run main.go Paris
//  go run main.go Moscow

func main() {
	city := os.Args[1]

	switchCaseIntro(city)
}

func switchCaseIntro(city string) {
	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	case "Moscow", "St. Petersburg":
		fmt.Println("Russia")
	default:
		fmt.Println("Where?")
	}
}
