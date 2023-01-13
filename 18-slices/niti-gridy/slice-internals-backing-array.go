package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"sort"
)

func main() {
	withABackingArr()
	fmt.Println("\n#####################")
	withACopyFunc()
}

func withACopyFunc() {
	grades := []float64{40, 10, 20, 50, 60, 70} // #4

	// #5: let's break the connection and copy orig in new arr
	var newGrades []float64
	newGrades = append(newGrades, grades...)
	// shortcut: newGrades := append([]float64(nil), grades...)

	front := newGrades[:3]
	sort.Float64s(front)

	// #7: new slices look at the same backing array
	// front, front2, front3, newGrades, they all have the same backing array
	front2 := front[:3]
	front3 := front

	s.PrintBacking = true          // #1
	s.MaxPerLine = 7               // #1
	s.Show("grades", grades[:])    // #1
	s.Show("newGrades", newGrades) // #5
	s.Show("front2", front2)       // #7
	s.Show("front3", front3)       // #7
}

func withABackingArr() {
	// #1: arrays and non-empty slice literals create an array.
	// For the arrays, it's explicit, but for the slices,
	// it's done implicitly, behind the scenes.

	grades := [...]float64{40, 10, 20, 50, 60, 70} // #1

	// #2: cheap: slicing doesn't allocate new memory (array).
	front := grades[:3]

	// #3: sort its first segment
	sort.Float64s(front)

	s.PrintBacking = true       // #1
	s.MaxPerLine = 7            // #1
	s.Show("grades", grades[:]) // #1
	s.Show("front", front)      // #2
}
