package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//games := loadGames()
	games := loadGamesFrom(data)
	gamesByID := indexById(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		menu()

		if !in.Scan() || !runCmd(in.Text(), games, gamesByID) {
			break
		}
	}
}
