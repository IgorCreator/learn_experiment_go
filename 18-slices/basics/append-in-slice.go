package main

import s "github.com/inancgumus/prettyslice"

func main() {
	sliceWithNums()
	sliceWithString()

}

func sliceWithString() {
	var todo []string

	todo = append(todo, "sing")

	// you can only append elements with the same element type of the slice
	// todoArr = append(todoArr, 42)

	todo = append(todo, "run")

	// append is a variadic function, so you can append multiple elements
	todo = append(todo, "code", "play")

	// you can also append a slice to another slice using ellipsis: ...
	tomorrow := []string{"see mom", "learn go"}
	todo = append(todo, tomorrow...)

	s.Show("todo", todo)
}

func sliceWithNums() {
	nums := []int{5, 6, 7}
	nums = append(nums, 7, 1, 9)
	s.Show("nums", nums)

	tens := []int{10, 20}
	nums = append(nums, tens...)
	s.Show("nums", nums)
}
