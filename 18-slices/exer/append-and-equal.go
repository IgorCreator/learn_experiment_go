package main

import (
	"bytes"
	"fmt"
	s "github.com/inancgumus/prettyslice"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, png...)

	s.Show("png", png)
	s.Show("header", header)
	fmt.Printf("Slices are: %t\n", bytes.Equal(header, png))
}
