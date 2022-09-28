package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	calc()
	bye()
}

func calc() {
	if 5 > 1 {
		fmt.Printf("5 ")
		fmt.Printf("bigger")
		fmt.Println(" then 1")
	}
}
