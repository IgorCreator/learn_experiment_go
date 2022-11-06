package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run scope.go
// go run scope.go hoho
// go run scope.go 3

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// a, n and err are available here
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// all the variables (names) are available here
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// a, n and err are not available here
	// they belong to the if statement
}
