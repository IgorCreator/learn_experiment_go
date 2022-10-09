package main

import "fmt"

func main() {
	convertSpeed()
	fmt.Println("------------------")
	ex1()
	fmt.Println("------------------")
	ex2()
	fmt.Println("------------------")
	ex3()
	fmt.Println("------------------")
	ex4AreaManipulation()
	fmt.Println("------------------")
}

func convertSpeed() {
	speed := 100 // int
	force := 2.5 // float64

	fmt.Printf("speed     : %T\n", speed)
	fmt.Printf("conversion: %T\n", float64(speed))
	fmt.Printf("expression: %T\n", float64(speed)*force)

	speed = int(float64(speed) * force)
	fmt.Printf("Rersult: %d (type: %T)\n", speed, speed)
}

func ex4AreaManipulation() {
	area := 10.5
	fmt.Println(area / 2)

	div := 2
	//fmt.Print(sq / div) //can't divide different type of values.
	fmt.Println(area / float64(div))
}

func ex3() {
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE: 103 - Wrong
	fmt.Println(min)
	fmt.Println(int8(max))
	fmt.Println(int8(max) + min)

	// Solution: res = 1127
	fmt.Println(max + int16(min))
	// EXPLANATION
	//
	// `int8(max)` destroys the information of max
	// It reduces it to 127
	// Which is the maximum value of int8
	//
	// Correct conversion is int16(min)
	// Because, int16 > int8
	// When you do so, min doesn't lose information
}

func ex2() {
	a, b := 10, 5.5
	fmt.Println(a, b)
	a = int(b)
	fmt.Println(a, b)
	fmt.Println(float64(a) + b)
}
func ex1() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	age := 2
	fmt.Println(7.5 + float64(age))
	years := 3. // No need for conversion as it already float64
	fmt.Println(7.5 + years)
}
