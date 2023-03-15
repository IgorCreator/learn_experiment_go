package main

// cd 25-oop/nonstructs
// go run .
// time go run .

func main() {
	var (
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	var items []*game
	items = append(items, &minecraft, &tetris)

	// you can attach methods to a compatible type on the fly:
	// items -> []*game
	// list  -> []*game
	my := list(items)

	// you can call methods even on a nil value
	my.print()
}
