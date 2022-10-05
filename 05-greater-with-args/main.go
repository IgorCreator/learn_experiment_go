package main

import (
	"fmt"
	"os"
)

//to run
//  go build -o greeter
// 	./greeter hi hello

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path: ", os.Args[0])

	if len(os.Args) == 2 {
		fmt.Println("1st argument: ", os.Args[1])
	}
	if len(os.Args) == 3 {
		fmt.Println("1st argument: ", os.Args[1])
		fmt.Println("2nd argument: ", os.Args[2])
	}

}
