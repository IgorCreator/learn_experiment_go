package main

import (
	"encoding/json"
	"fmt"
	i "go_test/25-oop/marshaler-interface/intefaces"
	"log"
)

// go run ./..

const data = `
[
 {
  "Title": "moby dick",
  "Price": 10,
  "Released": 118281600
 },
 {
  "Title": "odyssey",
  "Price": 15,
  "Released": 733622400
 },
 {
  "Title": "hobbit",
  "Price": 25,
  "Released": -62135596800
 }
]
`

func main() {
	decoder(data)
	encoder()
}

func decoder(data string) {
	var l i.List
	if err := json.Unmarshal([]byte(data), &l); err != nil {
		log.Fatal(err)
	}
	fmt.Print(l)
}

func encoder() {
	l := i.List{
		{Title: "moby dick", Price: 10, Released: i.ToTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: i.ToTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}

	data, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
