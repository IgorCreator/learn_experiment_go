package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run leap-year.go 2000

func main() {
	solLeapYear()

	lessonSolLeapYear()
}

func lessonSolLeapYear() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

func solLeapYear() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
	} else if year, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
	} else if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
