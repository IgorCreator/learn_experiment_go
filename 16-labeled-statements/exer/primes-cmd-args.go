package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run primes-cmd-args.go 2 4 6 7 a 9 c 11 x 12 13
// go run primes-cmd-args.go 1 2 3 5 7 A B C

func main() {
	nums := os.Args[1:]

	for _, num := range nums {
		if n, err := strconv.Atoi(num); err == nil {
			if isPrime(n) {
				fmt.Printf("%d ", n)
			}
		}
	}
	fmt.Println()
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}

	for i, w := 5, 2; i*i <= n; {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true
}
