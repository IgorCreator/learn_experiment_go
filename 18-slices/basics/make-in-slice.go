package main

import (
	s "github.com/inancgumus/prettyslice"
	"strings"
)

// Q: 	why should I use the make function?

// A: 	when a slice needs to grow and if there isn't enough capacity,
// 		the append function allocates a new back injury. It's EXPENSIVE Operation.
// 		So you can prevent it by allocating a large enough backing slice
// 		So when you grow slice, it won't allocate a new way if you define its capacity wisely,

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	basicsOfMakeFunc()
	effientProofUsageOfMake()
}

func effientProofUsageOfMake() {
	tasks := []string{"jump", "run", "read"}
	upTasks := make([]string, 0, len(tasks))
	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task)) // we won't create new slice. It is efficient
		s.Show("upTasks", upTasks)
	}
}

func basicsOfMakeFunc() {
	sl := make([]int, 3, 5)
	s.Show("init slice", sl)
	sl = append(sl, 42) // Will append as 4 el after all zero elements
	sl[0] = 10          // Will exchange 1st elements
	s.Show("slice", sl)
}
