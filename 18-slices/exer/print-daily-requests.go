package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"strings"
)

func main() {
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250, // 1st day: 1350 requests
		200, 400, 50, // 2nd day: 650 requests
		900, 800, 600, // 3rd day: 2300 requests
		750, 250, 100, // 4th day: 1100 requests
		// grand total: 5400 requests
	}

	// reqs := []int{
	// 	      500, 600, 250,
	// 	      200, 400, 50,
	// 	      900, 800, 600,
	// 	      750, 250, 100,
	// 	      150, 654, 235,
	// 	      320, 534, 765,
	// 	      121, 876, 285,
	// 	      543, 642,
	// 	      // the last element is missing (your code should be able to handle this)
	// }

	// ================================================
	// #1: Make a new slice with the exact size needed.

	sizeD := len(reqs) / N // you need to find the size.
	if len(reqs)%N != 0 {
		sizeD++
	}
	daily := make([][]int, 0, sizeD+1)
	s.Show("daily", daily)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	//counter := 1
	//var day []int
	//nDay := 0
	//for _, req := range reqs {
	//	if counter == N {
	//		daily[nDay] = append([]int{}, day...)
	//		counter = 0
	//		nDay++
	//	}
	//	day = append(day, req)
	//	counter++
	//}

	for N <= len(reqs) {
		daily = append(daily, reqs[:N]) // append the daily requests
		reqs = reqs[N:]                 // move the slice pointer for the next day
	}

	// add the residual data
	if len(reqs) > 0 {
		daily = append(daily, reqs)
	}

	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	var grand int
	for i, day := range daily {
		var sum int
		for i := 0; i < len(day); i++ {
			sum = sum + day[i]
		}
		fmt.Printf("%-10d%-10d\n", i, sum)
		grand += sum
	}
	fmt.Printf("%9s %-10d\n", "GRAND:", grand)
	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	// ...

	// ================================================
}
