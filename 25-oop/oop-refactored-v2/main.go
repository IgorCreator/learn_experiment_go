// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	i "go_test/25-oop/oop-refactored-v2/intefaces"
)

func main() {
	l := i.List{
		{Title: "moby dick", Price: 10, Released: i.ToTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: i.ToTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}

	l.Discount(.5)
	l.Print()
}
