package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const gameTurn = 5
const usagePicks = `
Provide [guess] [-v|optional]

Please guess any number from 0 to 10
`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usagePicks)
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Printf(usagePicks)
		return
	}

	var verbose bool
	if len(os.Args) == 3 {
		if param := os.Args[2]; param != "-v" {
			fmt.Printf(usagePicks)
			return
		} else {
			verbose = true
		}
	}

	rand.Seed(time.Now().Unix())
	var balancer int
	if guess > 10 {
		balancer = guess / 4
	}

	for turn := 1; turn <= gameTurn+balancer; turn++ {
		rangeNum := 10
		if guess >= 10 {
			rangeNum = guess
		}
		n := rand.Intn(rangeNum) + 1

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
