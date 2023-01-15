package main

import (
	"fmt"
)

// -- Conclusion about append function:
// You will see, the append function doubles the capacity until the 1024 elements,
// after the growth is slower. The growth ratio is much slower, around 25 percent.
// This happens because growing a larger area is much more costly than growing a smaller area.
// Think about creating an array with hundreds, thousands of elements and copying the preview say it's
// 50,000 of elements to it so that palm function slows down the growth ratio.
// It assumes that you're not going to appoint a lot of elements onto the larger array growth again.

// -- Final conclusion:
//So in summary, you can absolutely trust the app and function.
//Don't try to optimize it unless necessary.

func main() {
	growthRationWithAppend()
}

func growthRationWithAppend() {
	var (
		nums   []int
		oldCap float64
	)

	// loop 10 million times
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}
}
