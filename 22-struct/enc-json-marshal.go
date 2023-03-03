package main

import (
	"encoding/json"
	"fmt"
)

//go run enc-json-marshal.go > user_out.txt

type Permissions map[string]bool

type user struct {
	Name        string      `json:"user_Name"`
	Pwd         string      `json:"-"`
	Permissions Permissions `json:"Permissions,omitempty"`
}

func main() {
	users := []user{ // #2
		{"inanc", "1234", nil},
		{"god", "42", Permissions{"admin": true}},
		{"devil", "66", Permissions{"write": true}},
		{"bob", "321", Permissions{"read": true}},
		{"tom", "321", Permissions{"read": true, "exec": true}},
	}

	jsonSt, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonSt)
}
