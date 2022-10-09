package main

import "fmt"

// SYNTAX ERROR:
// "non-declaration statement outside function body"
// safe := true

// Uuse the normal declaration at the package scope.
var safe = true

func printDeclarationExamples() {
	fmt.Println("------------------")
	fmt.Println(safe)

	shortVarRedeclaration()
	shortVsNormal()
	shortVsGroupDec()
	groupDeclaration()
}

func shortVsNormal() {
	// -----------------------------------------------------
	// if you know the initial value
	// -----------------------------------------------------

	// DON'T DO THIS:
	// var width, height = 100, 50

	// DO THIS (concise):
	// width, height := 100, 50

	// -----------------------------------------------------
	// redeclaration
	// -----------------------------------------------------

	// DON'T DO THIS:
	// width = 50
	// color := red

	// DO THIS (concise):
	// width, color := 50, "red"

	// -----------------------------------------------------
	// initialization-and-short-declaration
	// -----------------------------------------------------
	// var name string = "Carl"
	// var name = "Carl"
	name := "Carl"

	// var isScientist bool = true
	// var isScientist = true
	isScientist := true

	// var age int = 62
	// var age = 62
	age := 62

	// var degree float64 = 5.
	// var degree = 5.
	degree := 5.

	fmt.Println(name, isScientist, age, degree)

	// type inference also works for variables
	//
	// Go gets the type of the variable and assigns it
	//   to the newly declared variable
	//
	// The type of the name2 variable is `string` now
	name2 := name
	fmt.Println(name2)
}

func shortVarRedeclaration() {
	age, yourAge := 10, 20
	age, ratio := 42, 3.14
	fmt.Println(age, yourAge, ratio)
}

func shortVsGroupDec() {
	// -----------------------------------------------------
	// if you don't know the initial value
	// -----------------------------------------------------

	// DON'T DO THIS:
	// score := 0

	// DO THIS:
	// var score int

	// -----------------------------------------------------
	// group variables for readability
	// -----------------------------------------------------

	// var (
	// 	video    string

	// 	duration int
	// 	current  int
	// )
}

func groupDeclaration() {
	var (
		name   string
		age    int
		famous bool
	)

	// Example #1
	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	// Example #2
	name = "Somebody"
	age = 20
	famous = false

	fmt.Println(name, age, famous)

	// Example #3
	// EXERCISE: Why this doesn't work? Think about it.

	// name = 20
	// name = age

	// save the previous value of the variable
	// to a new variable
	var prevName string
	prevName = name

	// overwrite the value of the original variable
	// by assigning to it
	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name :", name)
}
