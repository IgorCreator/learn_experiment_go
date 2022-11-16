package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// go run infinity-loop.go
// Stop this app using CTRL+C keys

func main() {
	for {
		fmt.Printf("\r%s Please Wait. Processing....\n", printRandomChar())
		time.Sleep(time.Duration(250) * time.Millisecond)
	}
}

func printRandomChar() string {
	chars := strings.Fields("\\ | / -")
	randomIdx := rand.Intn(len(chars))
	return chars[randomIdx]
}
