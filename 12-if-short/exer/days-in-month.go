package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// go run days-in-month.go 2000

func main() {
	solDaysInMonth()
}

func solDaysInMonth() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	month := strings.ToLower(os.Args[1])
	currentTime := time.Now()
	year := currentTime.Year()

	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)
	days := 28

	days, done := daysInMonths(month, days, leap)
	if done {
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}

func daysInMonths(m string, days int, leap bool) (int, bool) {
	if m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", m)
		return 0, true
	}
	return days, false
}
