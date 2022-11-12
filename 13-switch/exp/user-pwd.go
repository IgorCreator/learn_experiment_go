package main

import (
	"fmt"
	"os"
)

// go run user-pwd.go jack 18
// go run user-pwd.go inanc 1879

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	switch u, p := args[1], args[2]; {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		fmt.Printf(accessOK, u)
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
