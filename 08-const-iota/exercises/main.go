package main

import "fmt"

func main() {
	monthsEx()
	seasons()
}

func seasons() {
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}

func monthsEx() {
	const max = 12

	const (
		_ = (max - iota)
		Nov
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)
}
