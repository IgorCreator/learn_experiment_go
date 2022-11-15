package main

import (
	"fmt"
	"os"
	"strconv"
)

//  go run sum-common.go 11 100

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide min and max")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Please provide num")
	}
	if max < min {
		fmt.Println("Min can't be bigger then max")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i
		fmt.Printf("%v", i)
		if i != max {
			fmt.Printf(" + ")
		}
	}
	fmt.Println(" =", sum)
}
