package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run 01-atoi-error-handling.go 34
// go run 01-atoi-error-handling.go haha

func main() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// happy path (err is nil)
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}
