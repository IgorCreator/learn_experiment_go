package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go run 02-challenge-feet-to-meters.go 34

const usage = `
Feet to Meters
--------------
This program converts feet into meters.
Usage:
feet [feetsToConvert]`

func main() {
	//mineSolution()

	correctSolution()
}

func mineSolution() {
	feet := os.Args[1]

	n, err := strconv.Atoi(feet)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	meters := float64(n) * 0.3048
	fmt.Printf("SUCCESS: Converted %d feet into meters %.2f\n", n, meters)
}

func correctSolution() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: '%s' is not a number.\n", arg)
		return
	}

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
