package main

import (
	"encoding/json"
	"fmt"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

func loadGamesFrom(data string) []game {
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
		return nil
	}

	// load the data into usual game values
	return loadGamesWithLoop(decoded)
}

func loadGamesWithLoop(decoded []jsonGame) (games []game) {
	for _, dg := range decoded {
		games = addGame(games, newGame(dg.ID, dg.Price, dg.Name, dg.Genre))
	}
	return
}

func cmdSave(games []game) bool {
	var encodable []jsonGame
	for _, g := range games {
		encodable = addJsonGame(encodable, newJsonGame(g.id, g.price, g.name, g.genre))
	}

	out, err := json.MarshalIndent(encodable, "", "\t")
	if err != nil {
		fmt.Println("Sorry:", err)
		return true
	}

	fmt.Println(string(out))
	return false
}

func addJsonGame(games []jsonGame, g jsonGame) []jsonGame {
	return append(games, g)
}

func newJsonGame(id, price int, name, genre string) jsonGame {
	return jsonGame{id, name, genre, price}
}
