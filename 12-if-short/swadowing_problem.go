package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run swadowing_problem.go 43

func main() {
	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		// here use the main func's n and err
		// variables instead of its own

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}

	// if statement has calculated the result using
	// the main func's n variable
	//
	// so, that's why it prints the correct result here
	fmt.Println("n is", n)
}

func with_shadowing_problem() {
	// UNCOMMENT THIS TO SEE IT IN ACTION:
	// var n int

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// else if here shadows the main func's n
		// variable and it declares it own
		// with the same name: n

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// n here belongs to the main func
	// not to the if statement above

	// UNCOMMENT ALSO LINES BELOW TO SEE IT IN ACTION:
	// fmt.Printf("n is %d. 👻 👻 👻 - you've been shadowed ;-)\n", n)
}
