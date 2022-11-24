package main

import (
	"fmt"
	"strings"
)

func main() {
	// Set Literal with fixed length and ellipsis
	names := [3]string{
		"Einstein",
		"Tesla",
		"Shepard",
	}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [...]byte{'H', 'E', 'L', 'L', 'O'}

	// Set Array Elements
	var (
		ratios [1]float64
		alives [4]bool
		zero   [0]uint8
	)

	ratios[0] = 3.14145

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	_ = zero

	outputArrays(names, distances, data, ratios, alives, zero)
}

func outputArrays(names [3]string, distances [5]int, data [5]uint8, ratios [1]float64, alives [4]bool, zero [0]uint8) {
	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		// try the %c verb
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}
}

func printingArr(arrName string, arr []any) {

	fmt.Println(arrName)
	fmt.Println("====================")
	for i, v := range arr {
		fmt.Printf("%s[%d]: %v\n", arrName, i, v)
	}
}
