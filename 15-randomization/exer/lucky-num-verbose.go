package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurn = 11
const usageVerbose = `
Provide [guess] [-v|optional]

Please guess any number from 0 to 10
`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usageVerbose)
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Printf(usageVerbose)
		return
	}

	var verbose bool
	if len(os.Args) == 3 {
		if param := os.Args[2]; param != "-v" {
			fmt.Printf(usageVerbose)
			return
		} else {
			verbose = true
		}
	}

	rand.Seed(time.Now().Unix())

	for turn := 1; turn <= maxTurn; turn++ {
		n := rand.Intn(guess) + 1

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
