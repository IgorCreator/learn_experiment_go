package main

import "fmt"

func main() {
	var i int

loop:
	if i < 5 {
		fmt.Printf("%-2d - looping\n", i)
		i++
		goto loop
	}
	fmt.Println("done while pooping with `goto` label")
}
