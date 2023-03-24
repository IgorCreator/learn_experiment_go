// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package intefaces

import (
	"fmt"
)

type Product struct {
	Title    string
	Price    Money
	Released Timestamp
}

func (p *Product) String() string {
	return fmt.Sprintf("%-15s: %-8s (%s)", p.Title, p.Price, p.Released)
}

func (p *Product) Discount(ratio float64) {
	p.Price *= Money(1 - ratio)
}
