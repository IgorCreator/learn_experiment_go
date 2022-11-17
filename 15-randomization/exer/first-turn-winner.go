package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const max = 11
const usage = `
Provide [guess]

Please guess any number from 0 to 10
`

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(usage)
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Printf(usage)
		return
	}
	rand.Seed(time.Now().Unix())
	compNum := rand.Intn(max)

	if compNum == guess {
		var msg string
		switch rand.Intn(3) {
		case 0:
			msg = "You won from first try!"
		case 1:
			msg = "YOU'RE AWESOME. Guessed from first try!"
		default:
			msg = "You won!"
		}

		fmt.Printf("%s Num was: %d", msg, compNum)
		return
	} else {
		var msg string
		switch rand.Intn(3) {
		case 0:
			msg = "You guess is incorrect"
		case 1:
			msg = "YOU LOST. TRY AGAIN?"
		default:
			msg = "You loss!"
		}

		fmt.Printf("%s Num was: %d", msg, compNum)
		return
	}
}
