// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package intefaces

import (
	"strconv"
	"time"
)

// Timestamp stores, formats and automatically prints a Timestamp.
type Timestamp struct {
	// Timestamp anonymously embeds a time.
	// no need to convert a time value to a Timestamp value to use the methods of the time type.
	time.Time
}

func (ts Timestamp) String() string {
	if ts.IsZero() { // same as: ts.Time.IsZero()
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"
	return ts.Format(layout) // same as: ts.Time.Format(layout)
}

// toTimestamp returns a Timestamp value depending on the type of `v`.
func ToTimestamp(v interface{}) (ts Timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
