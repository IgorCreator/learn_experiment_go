package main

import "fmt"

func main() {
	fmt.Println("---------------------")
	printfVsPrintln()

	fmt.Println("---------------------")
	escSeq()

	fmt.Println("---------------------")
	prinfVarTypes()

	fmt.Println("---------------------")
	fmt.Println("Why use other verbs than %v? \nA: because: type-safety")
	printfGeneralUsage()

	fmt.Println("---------------------")
	printfExperiments()
}

func printfExperiments() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// swiss army knife %v verb
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// argument indexing - indexes start from 1
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)

	// why use other verbs than %v? because: type-safety
	// uncomment to see the warnings:
	//
	//fmt.Printf("Planet: %d\n", planet)
	//fmt.Printf("Distance: %s millions kms\n", distance)
	//fmt.Printf("Orbital Period: %t days\n", orbital)
	//fmt.Printf("Does %v has life? %f\n", planet, hasLife)

	// correct verbs:
	//fmt.Printf("Planet: %s\n", planet)
	//fmt.Printf("Distance: %d millions kms\n", distance)
	//fmt.Printf("Orbital Period: %f days\n", orbital)
	//fmt.Printf("Does %s has life? %t\n", planet, hasLife)
}

func printfGeneralUsage() {
	fmt.Printf("%v\n", "venus")
	fmt.Printf("%v\n", 4)
	fmt.Printf("%v\n", 98.56)
	fmt.Printf("%v\n", true)
}

func prinfVarTypes() {
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	fmt.Println("---------------------")
	num := 3.14
	fmt.Printf("%n\n", num)
	fmt.Printf("%f\n", num)
	fmt.Printf("%.2f\n", num)
	fmt.Printf("%9.2f\n", num)

	fmt.Println("---------------------")
	str := "Some string"
	fmt.Printf("%q\n", str) // prints the string in quoted-form like this ""
	fmt.Printf("%s\n", str) // prints the string without ""
}

func printfVsPrintln() {
	ops, ok, fail := 2350, 543, 433

	fmt.Println("total:", ops, "- success:", ok, "/", fail)
	fmt.Printf(
		"total: %d - success: %d / %d\n",
		ops, ok, fail,
	)
}

func escSeq() {
	fmt.Println("hihi")   // without newline
	fmt.Println("hi\nhi") // \n = escape sequence
	fmt.Println("hi\\n\"hi\"")
}
