package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	firstIssue()
	secondIssue()
}

func firstIssue() {
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

func secondIssue() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions", "extra cheese")

	fmt.Printf("toppings    : %s\n", toppings)
	fmt.Printf("pizza       : %s\n", pizza)
}
