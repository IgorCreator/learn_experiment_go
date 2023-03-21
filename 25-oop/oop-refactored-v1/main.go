// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	i "go_test/25-oop/oop-refactored-v1/intefaces"
	o "go_test/25-oop/oop-refactored-v1/specifc_obj"
)

func main() {

	store := i.List{
		&o.Book{i.Product{"moby dick", 10}, 118281600},
		&o.Book{i.Product{"odyssey", 15}, "733622400"},
		&o.Book{i.Product{"hobbit", 25}, nil},
		&o.Puzzle{i.Product{"rubik's cube", 5}},
		&o.Game{i.Product{"minecraft", 20}},
		&o.Game{i.Product{"tetris", 5}},
		&o.Toy{i.Product{"yoda", 150}},
	}

	store.Discount(.5)
	store.Print()

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := &book{product{"moby dick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)
}
