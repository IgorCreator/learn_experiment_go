package main

import (
	"fmt"
	"os"
	"strings"
)

// go run housing-prices-with-slicing.go 2 3

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// parse the data
	rows := strings.Split(data, "\n")
	cols := strings.Split(rows[0], separator)

	// default case: slice for all the columns
	from, to := 0, len(cols)

	// find the from and to positions depending on the command-line arguments
	args := os.Args[1:]
	for i, v := range cols {
		l := len(args)

		if l >= 1 && v == args[0] {
			from = i
		}

		if l == 2 && v == args[1] {
			to = i + 1 // +1 because the stopping index is a position
		}
	}

	// "from" cannot be greater than or equal to "to": reset invalid arg to 0
	if from >= to {
		from = 0
	}

	for i, row := range rows {
		cols := strings.Split(row, separator)

		// print only the requested columns
		for _, h := range cols[from:to] {
			fmt.Printf("%-15s", h)
		}
		fmt.Println()

		// print extra new line for the header
		if i == 0 {
			fmt.Println()
		}
	}
}
