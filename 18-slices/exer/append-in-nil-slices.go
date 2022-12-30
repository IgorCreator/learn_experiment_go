package main

import (
	s "github.com/inancgumus/prettyslice"
	"time"
)

func main() {
	var (
		pizza       []string
		graduations []int
		departures  []time.Time
		lights      []bool
	)

	pizza = append(pizza, []string{"pepperoni", "onions", "extra", "cheese"}...)
	graduations = append(graduations, []int{1998, 2005, 2018}...)

	now := time.Now()
	departures = append(departures, now, now.Add(time.Hour*24), now.Add(time.Hour*48))

	lights = append(lights, []bool{true, false, true}...)

	s.Show("pizza", pizza)
	s.Show("graduations", graduations)
	s.Show("departures", departures)
	s.Show("lights", lights)
}
