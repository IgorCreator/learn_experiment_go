package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide min and max")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || max < min {
		fmt.Println("Please provide  correct num")
		return
	}

	runWithWhileLoop(min, max)
	runWithIterativeForLoop(min, max)
}

func runWithWhileLoop(min int, max int) {
	var (
		sum int
		i   = min
	)
	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		fmt.Printf("%v", i)
		if i < max-1 {
			fmt.Printf(" + ")
		}

		sum += i
		i++
	}
	fmt.Println(" =", sum)
}

func runWithIterativeForLoop(min int, max int) {
	var sum int
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		fmt.Printf("%v", i)
		if i != max {
			fmt.Printf(" + ")
		}
	}
	fmt.Println(" =", sum)
}
