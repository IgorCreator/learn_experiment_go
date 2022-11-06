package main

import (
	"fmt"
	"os"
	"strconv"
)

//  go run main.go 3

func main() {
	if score, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("Error: ", err)
	} else {
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
}
