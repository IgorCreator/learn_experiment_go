package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// go run mood.go Tom

const usage = `
[your name] [positive|negative]
`

func main() {
	if len(os.Args) != 3 || (os.Args[2] != "positive" && os.Args[2] != "negative") {
		fmt.Printf(usage)
		return
	}

	mood := [2][3]string{
		{"good 👍", "happy 😀", "awesome 😎"},
		{"bad 👎", "sad 😞", "terrible 😩"},
	}

	rand.Seed(time.Now().UnixNano())
	if os.Args[2] == "positive" || os.Args[2] == "pos" {
		posMood := mood[0][rand.Intn(len(mood[0]))]
		fmt.Printf("%s feels %s", os.Args[1], posMood)
	} else {
		negMood := mood[1][rand.Intn(len(mood[1]))]
		fmt.Printf("%s feels %s", os.Args[1], negMood)
	}
}
