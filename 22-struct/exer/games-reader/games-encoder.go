package main

import (
	"encoding/json"
	"fmt"
)

//go run games-encoder.go > ./games.json

func main() {
	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	games := []jsonGame{
		{ID: 1, Name: "god of war", Price: 50, Genre: "action adventure"},
		{ID: 2, Name: "x-com 2", Price: 50, Genre: "strategy"},
		{ID: 3, Name: "minecraft", Price: 50, Genre: "sandbox"},
	}

	jsonSt, err := json.MarshalIndent(games, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonSt)
}
