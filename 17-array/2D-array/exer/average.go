package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please tell me numbers (maximum 5 numbers).\n")
		return
	}

	nums := os.Args[1:]
	var (
		sum float64
		qty int
	)
	fmt.Printf("Your numbers: [")
	for _, v := range nums {
		if num, err := strconv.ParseFloat(v, 64); err != nil {
			fmt.Printf("0 ")
		} else {
			fmt.Printf("%.0f ", num)
			sum += num
			qty++
		}
	}
	fmt.Printf("]\n")

	fmt.Printf("Average: %.0f\n", sum/float64(qty))
}
