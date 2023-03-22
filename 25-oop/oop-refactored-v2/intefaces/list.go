// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package intefaces

import (
	"strings"
)

type List []*Product

func (l List) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚."
	}

	var str strings.Builder
	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}

	return str.String()
}

func (l List) Discount(ratio float64) {
	for _, p := range l {
		p.Discount(ratio)
	}
}
