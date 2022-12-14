package main

import (
	"fmt"
)

//  go run main.go

func main() {
	switchCaseIntro("Paris")
	switchCaseIntro("Moscow")

	withFallThrough(142)
	withFallThrough(100)
	withFallThrough(1)
	withFallThrough(0)
	withFallThrough(-32)
	withFallThrough(-132)
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

func withFallThrough(i int) {
	switch {
	case i == 0:
		fmt.Print("zero")
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	case i < -100:
		fmt.Print("big ")
		fallthrough
	case i < 0:
		fmt.Print("negative ")
		fallthrough
	default:
		fmt.Print("number")
	}

	fmt.Println()
}
