package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

var games = map[int]game{
	1: {item{1, "god of war", 50}, "action adventure"},
	2: {item{2, "x-com 2", 50}, "strategy"},
	3: {item{3, "minecraft", 50}, "sandbox"},
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Please enter \"list\", id of the game or \"quit\"\n")

		for input.Scan() {
			command := strings.ToLower(input.Text())

			if id, err := strconv.Atoi(command); err == nil {
				retieveSpecificGame(id, games)
			} else if command == "list" || command == "l" {
				listAllGames(games)
			} else if command == "quit" || command == "q" {
				fmt.Printf("Bye\n")
				return
			} else {
				fmt.Printf("Invalid command!!!\n")
				continue
			}
		}
	}
}

func retieveSpecificGame(id int, games map[int]game) {
	if g, ok := games[id]; ok {
		fmt.Printf("%3s %-10s %-10s %-20s\n", "id", "name", "price", "genre")
		fmt.Printf("%3d %-10s %-10d %-20s\n", g.id, g.name, g.price, g.genre)
		return
	}

	fmt.Printf("This game not in our Database\n")
}

func listAllGames(games map[int]game) {
	fmt.Printf("%3s %-10s %-10s %-20s\n", "id", "name", "price", "genre")
	for _, g := range games {
		fmt.Printf("%3d %-10s %-10d %-20s\n", g.id, g.name, g.price, g.genre)
	}
}
