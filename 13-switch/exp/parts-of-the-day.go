package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Current time is: %s\n", time.Now())

	switch t := time.Now().Hour(); {
	case t >= 6 && t < 12:
		fmt.Println("Good morning")
	case t >= 12 && t < 18:
		fmt.Println("Good afternoon")
	case t >= 18 && t < 22:
		fmt.Println("Good evening")
	default:
		//case t >= 22 && t <= 5:
		fmt.Println("Good night")
	}
}
