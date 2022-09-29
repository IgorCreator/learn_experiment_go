package main

import (
	"fmt"
	"go_test/03-packages-and-scopes/custom_pack"
)

// package scope
const ok = true

func main() {
	custom_pack.Hi()
	custom_pack.Bye()

	// test and ok are visible here
	var test = "Test string."
	fmt.Println(test, ok)

	fmt.Println("------------")
	trigger_sound()
}

func methodInExecutablePackage() {
	fmt.Println("This line in Executable Package")
}
