package main

import (
	"fmt"
	"path"
)

// package scope
const ok = true

func main() {
	fmt.Println("-----------------")
	visibilityOfVar()

	fmt.Println("-----------------")
	initValuesInGo()

	fmt.Println("-----------------")
	typeConversionEx()

	fmt.Println("-----------------")
	pathSplit()
	pathSplitWithShortDeclaration()
}

func initValuesInGo() {
	shortDeclaration := "only local declaration can be short"
	fmt.Println(shortDeclaration)

	var (
		gear  int
		heat  float64
		off   bool
		brand string
	)
	var speed, velocity int
	var securityOn = true

	fmt.Println(gear)
	fmt.Println(speed, velocity)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Println(brand)
	fmt.Printf("%q\n", brand)
	fmt.Println(securityOn)
}

func pathSplit() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println(dir)
	fmt.Println(file)
}

func pathSplitWithShortDeclaration() {
	_, file := path.Split("css/main.css")
	fmt.Println(file)
}

func typeConversionEx() {
	speed := 100
	force := 2.5

	velocity := force * float64(speed) * 1.1
	fmt.Println(velocity)

	speed = int(force) * speed
	fmt.Println(speed)
	speed = int(force * float64(speed))
	fmt.Println(speed)
}
