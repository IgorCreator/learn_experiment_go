package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run 21-buffers.go < shakespeare.txt

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fmt.Printf("%s\n", in.Text())
	}
}
