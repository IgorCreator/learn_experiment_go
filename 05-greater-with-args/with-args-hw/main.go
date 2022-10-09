package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args)
	fmt.Printf("There are %d argument(s) sent in program.\n", count)

	printThePath()
	printFirstArg()
	greetArgs()
}

func greetArgs() {
	if len(os.Args) < 4 {
		return
	}

	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")
}

func printFirstArg() {
	if len(os.Args) >= 2 {
		fmt.Printf("Your name: %s\n", os.Args[1])
	}
}

func printThePath() {
	fmt.Printf("Path of program: %s\n", os.Args[0])
}
