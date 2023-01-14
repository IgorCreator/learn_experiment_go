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
	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		c := float64(cap(ages))
		if c != oldCap {
			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n",
				len(ages), c, c/oldCap)
		}
		oldCap = c
	}
}
