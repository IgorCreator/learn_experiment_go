// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package intefaces

import "fmt"

type Printer interface {
	Print()
	Discount(ratio float64)
}

type List []Printer

func (l List) Print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.Print()
	}
}

func (l List) Discount(ratio float64) {
	for _, it := range l {
		it.Discount(ratio)
	}
}
