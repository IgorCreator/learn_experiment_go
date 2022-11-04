package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.Itoa(42)
	fmt.Println(str)

	num, err := strconv.Atoi("Hi")
	printResultAtoi(err, num)

	age, err := strconv.Atoi("1000")
	printResultAtoi(err, age)
}

func printResultAtoi(err error, num int) {
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(num)
	}
}
