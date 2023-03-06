package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type GameObj struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Genre string `json:"genre,omitempty"`
	Price int    `json:"price,omitempty"`
}

//go run games-decoder.go < ./games.json

func main() {
	in := bufio.NewScanner(os.Stdin)

	var input []byte
	var games []GameObj
	for in.Scan() {
		input = append(input, in.Bytes()...)
	}

	err := json.Unmarshal(input, &games)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	} else {
		listAllGames(games)
	}
}

func listAllGames(games []GameObj) {
	fmt.Printf("%3s %-10s %-10s %-20s\n", "id", "name", "price", "genre")
	for _, g := range games {
		fmt.Printf("%3d %-10s %-10d %-20s\n", g.Id, g.Name, g.Price, g.Genre)
	}
}
