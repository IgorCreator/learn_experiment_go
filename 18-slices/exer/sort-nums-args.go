package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
	"os"
	"sort"
	"strconv"
)

// go run sort-nums-args.go  4 6 1 3 0 9 2
// go run sort-nums-args.go  4 a 1 c d 9

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	var res []int
	for i := range args {
		num, err := strconv.Atoi(args[i])
		if err == nil {
			res = append(res, num)
		}
	}

	sort.Ints(res)
	s.Show("Result", res)
}
