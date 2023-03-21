// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package specifc_obj

import (
	"fmt"
	"go_test/25-oop/oop-refactored-v1/intefaces"
	"strconv"
	"time"
)

type Book struct {
	intefaces.Product
	Published interface{}
}

func (b *Book) Print() {
	// the book can also call the embedded product's print method
	// if it wants to, as in here:
	b.Product.Print()

	p := format(b.Published)
	fmt.Printf("\t - (%v)\n", p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}

	const layout = "2006/01"
	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
