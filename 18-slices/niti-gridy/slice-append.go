package main

import (
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 8

	var nilArr []int
	s.Show("nillArray", nilArr)

	// #1 simple append
	ages := []int{35, 15}
	s.Show("ages", ages)
	ages = append(ages, 5)
	s.Show("ages", ages)
	ages = append(ages, 10)
	s.Show("ages", ages)

	// #2 Add in the middle
	ages = append(ages, ages[2:]...)
	s.Show("duplicate middle el", ages)
	ages = append(ages[:2], 7, 9)
	s.Show("change middle elements", ages)
	ages = ages[:6]
	s.Show("revel all el(backing arr)", ages)

	// #3 doubling backing array for slices
	seeDoublingWithAppend()
}

func seeDoublingWithAppend() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)

		time.Sleep(time.Second / 4)
	}
}
