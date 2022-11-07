package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run odd-even.go 43

func main() {
	mySolutionOddEven()

	lessonSolOddEven()
}

func lessonSolOddEven() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}
}

func mySolutionOddEven() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
	} else if num, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
	} else if num%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8.\n", num)
	} else if num%2 == 0 {
		fmt.Printf("%d is an even.\n", num)
	} else {
		fmt.Printf("%d is an odd.\n", num)

	}
}
