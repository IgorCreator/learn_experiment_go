package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// go run mood.go Tom

const usage = `
[your name]
`

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(usage)
		return
	}

	mood := [...]string{
		"good 👍",
		"bad 👎",
		"sad 😞",
		"happy 😀",
		"awesome 😎",
		"terrible 😩",
	}

	rand.Seed(time.Now().UnixNano())
	randMood := mood[rand.Intn(len(mood))]
	fmt.Printf("%s feels %s", os.Args[1], randMood)
}
