package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usageApp = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
This is why guess two numbers.
Wanna play?

[guess1] [guess2]
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usageApp, maxTurns)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Not a numbers.")
		return
	}

	if guess1 <= 0 || guess2 <= 0 {
		fmt.Println("Please pick a positive numbers.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n1 := rand.Intn(guess1) + 1
		n2 := rand.Intn(guess2) + 1

		if n1 != guess1 && n2 != guess2 {
			continue
		}

		if n1 == guess1 && n2 == guess2 && turn == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else if n1 == guess1 && n2 == guess2 {
			fmt.Println("🎉  YOU Guessed both nums!")
		} else if n1 == guess1 {
			fmt.Println("🎉  YOU Guessed first num!")
		} else if n2 == guess2 {
			fmt.Println("🎉  YOU Guessed second num!")
		}
		return
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
