package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	history := make(map[string]bool)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		command := strings.ToLower(input.Text())
		if history[command] {
			fmt.Printf("This already was typed. Quiting....\n")
			return
		}

		history[command] = true
	}
}
