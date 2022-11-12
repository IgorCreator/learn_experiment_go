package main

import (
	"fmt"
	"os"
	"strings"
)

// go run string-manipul.go lower 'OMG!'
// go run string-manipul.go upper 'omg!'
// go run string-manipul.go title "mr. charles darwin"

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Please send 2 args [command] [string]\n " +
			"Where [command]: lower, upper and title\n " +
			"And [string]: any string in \"\"")
		return
	}

	switch cmd := os.Args[1]; cmd {
	case "lower":
		fmt.Printf("%s\n", strings.ToLower(os.Args[2]))
	case "upper":
		fmt.Printf("%s\n", strings.ToUpper(os.Args[2]))
	case "title":
		fmt.Printf("%s\n", strings.Title(os.Args[2]))
	default:
		fmt.Printf("Unknown command: %q\n", cmd)
	}
}
