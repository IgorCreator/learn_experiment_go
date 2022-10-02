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

	fmt.Println("------------")
	// 1.test and ok are visible here
	var test = "Test string."
	fmt.Println(test, ok)

	fmt.Println("------------")
	//2. trigger_sound method visible in main package because it located in the same package
	trigger_sound()
	fmt.Println("------------")
	//3. Version method located in different package, so we need import and have capitalize
	fmt.Println(custom_pack.Version())
}

func methodInExecutablePackage() {
	fmt.Println("This line in Executable Package")
}
