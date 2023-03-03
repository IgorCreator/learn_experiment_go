package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// go run decode-json-unmarshal.go > user_out.txt
// go run decode-json-unmarshal.go < user_out.txt

type Employee struct {
	Name        string          `json:"user_Name"`
	Permissions map[string]bool `json:"Permissions,omitempty"`
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	var input []byte
	for in.Scan() {
		input = append(input, in.Bytes()...)
	}

	var empls []Employee
	err := json.Unmarshal(input, &empls)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("%v\n", empls)

	for _, employee := range empls {
		fmt.Print("+ ", employee.Name)

		switch p := employee.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		case p["read"]:
			fmt.Print(" can read.")
		case p["exec"]:
			fmt.Print(" can execute.")
		}
		fmt.Println()
	}
}
