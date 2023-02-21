package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel", // #5
	}

	// changing both maps dict and turkish
	//turkish := dict
	//turkish["good"] = "güzel"
	//dict["great"] = "kusursuz"

	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}
	fmt.Printf("english dict len: %d\nturkish dict len: %d\n", len(dict), len(turkish))
	fmt.Printf("english: %q\nturkish: %q\n", dict, turkish) // #4

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.\n", query)
}
