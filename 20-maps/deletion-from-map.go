package main

import "fmt"

func main() {
	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel", // #5
	}
	fmt.Printf("english: %q\n", dict)

	delete(dict, "awesome") // #6
	delete(dict, "awesome") // #7: no-op
	delete(dict, "notexisting")
	fmt.Printf("english: %q\n", dict)

	// dict = nil // map not deleted
	for k := range dict {
		delete(dict, k)
	}
	fmt.Printf("english: %q\n", dict)
}
