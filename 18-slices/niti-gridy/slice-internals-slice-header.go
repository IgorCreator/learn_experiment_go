package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

type collection []string

func main() {
	// SliceHeader lives here:
	// https://golang.org/src/runtime/slice.go

	s.PrintElementAddr = true
	data := collection{"slices", "are", "awesome", "period", "!!"}

	changeColl(data)

	s.Show("main's data", data)
	fmt.Printf("main's data slice's header: %p\n", &data)

	// ----------------------------------------------------------------------
	// #4
	array := [5]string{"Arrays", "are", "awesome", "period", "!!" /* #5 */}
	changeArr(array)
	s.Show("main's data", array)

	// array's size depends on its elements
	fmt.Printf("array's size: %d bytes.\n", unsafe.Sizeof(array))
	// slice's size is always fixed: 24 bytes (on a 64-bit system) — slice value = slice header
	fmt.Printf("slice's size: %d bytes.\n", unsafe.Sizeof(data))
}

// #1
// passed value will be copied in the function
func changeColl(data collection) {
	// data is a new variable inside the function:
	// var data collection

	data[2] = "brilliant!"

	s.Show("change's data", data)
	fmt.Printf("change's data slice's header: %p\n", &data) // #3
}

func changeArr(data [5]string) {
	// arrays do copies of originals

	data[2] = "brilliant!"

	s.Show("change's data", data)
	fmt.Printf("change's data slice's header: %p\n", &data) // #3
}
