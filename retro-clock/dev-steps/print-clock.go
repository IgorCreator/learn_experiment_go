package main

import (
	"fmt"
	"strconv"
	"time"
)

type placeholder [5]string

func main() {

	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()
	fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

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

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// 1st Sol
	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		separator,
		digits[min/10], digits[min%10],
		separator,
		digits[sec/10], digits[sec%10],
	}

	// 2nd Sol
	//clock := [...]placeholder{
	//	digits[extractSpecificDig(0, hour)],
	//	digits[extractSpecificDig(1, hour)],
	//	separator,
	//	digits[extractSpecificDig(0, min)],
	//	digits[extractSpecificDig(1, min)],
	//	separator,
	//	digits[extractSpecificDig(0, sec)],
	//	digits[extractSpecificDig(1, sec)],
	//}

	for line := range clock[0] {
		for _, num := range clock {
			fmt.Printf("%s  ", num[line])
		}
		fmt.Printf("\n")
	}
}

func extractSpecificDig(pos int, num int) int {
	slice := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(slice[pos])) // Doesn't work if one digit
	fmt.Printf("%d\n", res)
	return res
}
