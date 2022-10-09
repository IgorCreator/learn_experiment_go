package main

import "fmt"

func runMultiplier(n float64, m int) {
	x := n * float64(m)
	fmt.Printf("Result: %f\n", x)
}

func findPerimeter() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2 * (width + height)
	fmt.Printf("Result: %d\n", perimeter)
}
