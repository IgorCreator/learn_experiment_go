package main

import (
	"fmt"
	i "go_test/25-oop/oop-refactored-v2/intefaces"
	"sort"
)

// go run ...
func main() {
	l := i.List{
		{Title: "moby dick", Price: 10, Released: i.ToTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: i.ToTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}
	l.Discount(.5)

	//sort.Sort(l)
	//sort.Sort(sort.Reverse(l))
	//sort.Sort(i.ByReleasedDate(l))
	sort.Sort(sort.Reverse(i.ByReleasedDate(l)))

	fmt.Printf("%s", l)
}
