package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please give me numbers\n")
		return
	}

	var nums []float64
	var counter int
	for _, v := range os.Args[1:] {
		if num, err := strconv.ParseFloat(v, 64); err != nil {
			fmt.Printf("Skipped not number: %v", num)
			continue
		} else {
			nums = append(nums, num)
			counter++
		}
	}

	sort.Float64s(nums)
	fmt.Printf("Sorted arr: %v", nums)
}
