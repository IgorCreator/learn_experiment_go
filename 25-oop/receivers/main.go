package main

// cd 25-oop/receivers
// go run .
// time go run .

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// sends a pointer of minecraft to the discount method
	// same as: (&minecraft).discount(.1)
	minecraft.discount(.1)

	mobydick.print()
	minecraft.print()
	tetris.print()

	// creates a variable that occupies 200mb on memory
	var h huge
	for i := 0; i < 10; i++ {
		h.addrWithPtr() // much faster and appropriet impl
		//h.addrWithObjCopy()   // very slow version
	}
}
