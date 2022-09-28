package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello" + "!!!")
	calc()
	fmt.Printf("Logical CPUs: %s", runtime.NumCPU())
	bye()
}

func calc() {
	if 5 > 1 {
		fmt.Printf("5 ")
		fmt.Printf("bigger")
		fmt.Println(" then 1")
	}
}
