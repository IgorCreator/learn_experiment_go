package main

import (
	"fmt"
	"os"
	"strconv"
)

//  go build -o if-clause
// 	./if-clause 10

func main() {
	score, _ := strconv.Atoi(os.Args[1])

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else if score == 2 {
		fmt.Println("meh...")
	} else {
		fmt.Println("low")
	}
}
