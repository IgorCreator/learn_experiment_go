package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	// it stops the garbage collector: prevents cleaning up the memory.
	debug.SetGCPercent(-1)
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	var arr [size]int
	report("after declaring an array")

	// 3. copy the array to a new array.
	cpArr := arr
	report("after copying the array")

	// 5. pass the array to the passArray function
	passArray(cpArr)

	// 6. convert one of the arrays to a slice
	sl := arr[:]
	report("after convertion")
	// 7. slice only the first 1000 elements of the array
	sl_1 := sl[:1000]
	report("after 1st slicing")
	// 8. slice only the elements of the array between 1000 and 10000
	sl_2 := sl[1000:10000]
	report("after 2nd slicing")

	// 10. pass the one of the slices to the passSlice function
	passSlice(sl_1)
	passSlice(sl_2)

	// 11. print the sizes of the arrays and slices
	fmt.Printf("Array's size  : %d bytes.\n", unsafe.Sizeof(arr))
	fmt.Printf("Array2's size : %d bytes.\n", unsafe.Sizeof(cpArr))
	fmt.Printf("Slice1's size : %d bytes.\n", unsafe.Sizeof(sl))
	fmt.Printf("Slice2's size : %d bytes.\n", unsafe.Sizeof(sl_1))
	fmt.Printf("Slice3's size : %d bytes.\n", unsafe.Sizeof(sl_2))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
