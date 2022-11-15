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
	fmt.Println("---------------------------------------------")
	combinedSolution(3, "/")
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

func combinedSolution(size int, op string) {
	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var res int

			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}
