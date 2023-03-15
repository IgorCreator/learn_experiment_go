package main

import "fmt"

type game struct {
	title string
	price float64
}

// be consistent: use pointer everywhere if another method in the same type has a pointer receiver
func (g *game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

// be consistent: use pointer everywhere if another method in the same type has a pointer receiver
func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
}
