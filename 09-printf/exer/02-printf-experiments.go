package main

import "fmt"

func falseClaims() {
	tf := false
	fmt.Printf("These are %t claims.\n", tf)
}

func temp() {
	fmt.Printf("Temperature is %.1f degrees.\n", 29.5)
}

func quotes() {
	fmt.Printf("%q\n", "hello world")
}

func typePrint() {
	fmt.Printf("Type of %d is %[1]T\n", 3)
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
	fmt.Printf("Type of %s is %[1]T\n", "hello")
	fmt.Printf("Type of %t is %[1]T\n", true)
}
