package main

import "fmt"

func swapper(a string, b string) {
	a, b = b, a

	fmt.Println(a, b)
}
