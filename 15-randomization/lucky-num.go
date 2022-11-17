package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please submit number for generation")
	} else if num, err := strconv.Atoi(os.Args[1]); err != nil || num < 0 {
		fmt.Println("Please provide actual positive int num")
		return
	} else {
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < num; i++ {
			randomIdx := rand.Intn(num + 1)
			fmt.Printf("%d ", randomIdx)
		}
		fmt.Println()
	}
}
