package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"runtime"
)

func main() {
	// reports the initial memory usage
	Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := Read()

	// -----------------------------------------------------
	last10 := make([]int, 10)
	copy(last10, millions[len(millions)-10:])

	//a)millions = last10
	// or
	//b)millions = append([]int(nil), millions[len(millions)-10:]...)
	millions = last10
	fmt.Printf("\nLast 10 elements: %d\n\n", last10)
	// -----------------------------------------------------

	Report()

	// don't worry about this code.
	fmt.Fprintln(ioutil.Discard, millions[0])
}

// Read returns a huge slice (allocates ~65 MB of memory)
func Read() []int {
	// 2 << 22 means 2^(22 + 1)
	// See this: https://en.wikipedia.org/wiki/Arithmetic_shift

	// Perm function returns a slice with random integers in it.
	// Here it returns a slice with random integers that contains
	// 8,388,608 elements. One int value is 8 bytes.
	// So: 8,388,608 * 8 = ~65MB
	return rand.Perm(2 << 22)
}

// Report cleans the memory and prints the current memory usage
// Don't worry about this code. You don't need to understand it.
//
// However, if you're curious, read on.
//
// The following code runs the garbage collector to clean
// up the allocated resources, and then it reads the current
// memory statistics into the m variable.
func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
