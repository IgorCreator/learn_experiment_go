package main

import (
	s "github.com/inancgumus/prettyslice"
	"strconv"
)

func main() {
	nums := []string{"2", "4", "6", "1", "3", "5"}
	l := len(nums)

	var even, odd []int
	for _, num := range nums {
		i, _ := strconv.Atoi(num)
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}

	s.Show("Even", even)
	s.Show("odd", odd)

	s.Show("middle", nums[l/2-1:l/2+1])
	s.Show("two first", nums[:2])
	s.Show("two last", nums[l-2:])

	s.Show("last even", even[len(even)-1:])
	s.Show("last odd", odd[len(odd)-2:])
}
