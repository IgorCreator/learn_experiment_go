package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("[english word] -> [turkish word]")

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me english word")
		return
	}
	query := args[0]

	// #1: Empty Map Literal
	// dict := map[string]string{}

	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
		// 42: "forty two",  // can't use, diff types
		// "forty two": 42,  // can't use, diff types
	}

	dict["up"] = "yukarı"  // adds a new pair
	dict["down"] = "aşağı" // adds a new pair
	dict["good"] = "iyi"   // #5: overwrites the value at the key: "good"
	dict["mistake"] = ""   // #8: a key with a zero-value

	if value, existsInMap := dict[query]; existsInMap {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found.\n", query)
	fmt.Printf("Dict contains %d records\n", len(dict))
}
