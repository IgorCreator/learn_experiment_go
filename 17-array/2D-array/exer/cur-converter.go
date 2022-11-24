package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)

func main() {
	rates := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	if len(os.Args) != 2 {
		fmt.Printf("Please privovide amount of USD to conver\n")
		return
	}

	if amount, err := strconv.ParseFloat(os.Args[1], 32); err != nil {
		fmt.Println("Please provide amount in int or float format")
	} else {
		//forLoopThroughArr(rates, amount)

		fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
		fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP]*amount)
		fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY]*amount)

	}
}

func forLoopThroughArr(rates [3]float64, amount float64) {
	for i := range rates {
		fmt.Printf("%d USD is %g [currentcy: %s]\n", amount, amount*rates[i], rates[i])
	}
}
