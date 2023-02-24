package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go run grep-clone.go < ./../shakespeare.txt
// go run grep-clone.go come < ./../shakespeare.txt

func main() {
	//personalSol()
	optimizedSol()
}

func personalSol() {
	args := os.Args
	input := bufio.NewScanner(os.Stdin)

	if len(args) > 1 {
		query := strings.ToLower(args[1])

		for input.Scan() {
			scannedLine := strings.ToLower(input.Text())
			if strings.Contains(scannedLine, query) {
				fmt.Printf("%s\n", scannedLine)
			}
		}
	} else {
		for input.Scan() {
			fmt.Printf("%s\n", input.Text())
		}
	}
}

func optimizedSol() {
	in := bufio.NewScanner(os.Stdin)

	var pattern string
	if args := os.Args[1:]; len(args) == 1 {
		pattern = args[0]
	}

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}
}
