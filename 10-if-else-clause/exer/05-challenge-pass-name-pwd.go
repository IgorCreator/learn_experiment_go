package main

import (
	"fmt"
	"os"
)

//  go build -o args-app 05-arg-count.go
// 	./args-app
// 	./args-app albert
// 	./args-app hacker 42
// 	./args-app jack 6475
// 	./args-app jack 1888
// 	./args-app inanc 1879

var dict = map[string]string{
	"jack":  "1888",
	"inanc": "1879",
}

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	username, pwd := args[1], args[2]
	if val, ok := dict[username]; ok {
		if val == pwd {
			fmt.Printf(accessOK, username)
		} else {
			fmt.Printf(errPwd, username)
		}
	} else {
		fmt.Printf(errUser, username)
	}
}
