package main

import "fmt"

type huge struct {
	// about ~200mb
	games [10_000_000]game
}

// only copies a single pointer.
func (h *huge) addrWithPtr() {
	fmt.Printf("%p\n", h)
}

// each time it is called,
// the original value will be copied.
// calling addr() x 10 times = ~2 GB.
func (h huge) addrWithObjCopy() {
	fmt.Printf("%p\n", &h)
}
