package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type placeholder [5]string

const sc = "\u001B7"
const rc = "\u001B8"

// go run clock.go templates.go

func main() {
	digits := FormDigits()

	hide := false
	for {
		clearScreen()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			GetSeparator(hide),
			digits[min/10], digits[min%10],
			GetSeparator(hide),
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for _, num := range clock {
				fmt.Printf("%s  ", num[line])
			}
			fmt.Printf("\n")
		}

		time.Sleep(time.Second)
		hide = !hide
	}

}

func clearScreen() {
	fmt.Print(rc + sc)
	cmd := exec.Command("clear") //Unix example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
