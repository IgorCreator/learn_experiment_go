package main

import (
	"fmt"
)

func main() {
	fmt.Println("[english word] -> [turkish word]")

	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	dict["up"] = "yukarı"  // adds a new pair
	dict["down"] = "aşağı" // adds a new pair
	dict["good"] = "iyi"   // #5: overwrites the value at the key: "good"
	dict["mistake"] = ""   // #8: a key with a zero-value

	//#13: compare a map using its printed output
	copied := map[string]string{"up": "yukarı", "down": "aşağı",
		"mistake": "", "good": "iyi", "great": "harika",
		"perfect": "mükemmel"}
	first := fmt.Sprintf("%s", dict)
	second := fmt.Sprintf("%s", copied)
	if first == second {
		fmt.Println("Maps are equal")
	}

	//Printing a map (ordered output since Go 1.12)
	fmt.Printf("%#v\n", dict)
	//Not Efficient to pring all values. Use %v
	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v)
	}

}
