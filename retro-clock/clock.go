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

func main() {
	digits := formDigits()

	hide := false
	for {
		clearScreen()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			getSeparator(hide),
			digits[min/10], digits[min%10],
			getSeparator(hide),
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

func formDigits() [10]placeholder {
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	return digits
}

func getSeparator(hide bool) placeholder {
	if hide {
		return placeholder{
			"   ",
			" ░ ",
			"   ",
			" ░ ",
			"   ",
		}
	} else {
		return placeholder{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		}
	}
}

func clearScreen() {
	fmt.Print(rc + sc)
	cmd := exec.Command("clear") //Unix example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
