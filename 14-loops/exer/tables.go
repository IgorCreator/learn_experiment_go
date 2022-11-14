package main

import "fmt"

//  go run tables.go

func main() {
	fmt.Println("---------------------------------------------")
	multiplicationTable(5)
	fmt.Println("---------------------------------------------")
	additionTable(3)
	fmt.Println("---------------------------------------------")
	subtractionTable(4)
	fmt.Println("---------------------------------------------")
	divisionTable(6)
}

func multiplicationTable(max int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

func additionTable(max int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i+j)
		}
		fmt.Println()
	}
}

func subtractionTable(max int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i-j)
		}
		fmt.Println()
	}
}

func divisionTable(max int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%.2f", float64(i)/float64(j))
		}
		fmt.Println()
	}
}
