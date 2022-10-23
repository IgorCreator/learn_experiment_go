package main

import (
	"fmt"
)

func main() {
	resUrl := "https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3"
	fmt.Println("More about iota here:" + resUrl)
	weekDays()
	months()
	timezones()
}

func timezones() {
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println("Timezons: ", EST, MST, PST)
}

func weekDays() {
	const (
		mon = iota
		tue
		wend
		thru
		fri
		sat
		sun
	)

	fmt.Println(
		"Days: ",
		mon,
		tue,
		wend,
		thru,
		fri,
		sat,
		sun,
	)
}

func months() {
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
		Jul
		Aug
		Sep
		Oct
		Nov
		Dec
	)

	fmt.Println(
		"Months: ",
		Jan,
		Feb,
		Mar,
		Apr,
		May,
		Jun,
		Jul,
		Aug,
		Sep,
		Oct,
		Nov,
		Dec,
	)
}
