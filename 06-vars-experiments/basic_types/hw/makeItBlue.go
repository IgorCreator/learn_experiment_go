package main

import "fmt"

func runProgMakeItBlue() {
	fmt.Printf("EXERCISE: Make It Blue\n")
	color := "green"
	fmt.Printf("Color in the beginning: %s\n", color)
	color = "blue"
	fmt.Printf("Color in the end: %s\n", color)
}

func runVarsToVars(color string) {
	fmt.Printf("Color now: %s\n", color)
}
