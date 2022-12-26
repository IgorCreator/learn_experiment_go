package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	names := []string{"Einstein", "Shepard", "Tesla"}

	books := []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}
	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}
	// use the slicing expression to change the nums array to a slice below
	sort.Ints(nums[:])

	//Here: Use the strings.Join function to join the names
	fmt.Printf("%q\n", strings.Join(names, " "))

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
