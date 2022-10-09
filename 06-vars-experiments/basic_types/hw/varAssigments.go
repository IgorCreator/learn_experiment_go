package main

import "fmt"

func multiAssign() {
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")
}

func multiReturnWithDiscardingval() {
	_, b := multi()
	fmt.Println(b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}
