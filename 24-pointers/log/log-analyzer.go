// go run . < ./logs/log.txt
// go run . < ./logs/log_err_missing.txt
// go run . < ./logs/log_err_negative.txt
// go run . < ./logs/log_err_str.txt

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parsed := result{}
	p := newParser()

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.lines++

		err := parse(in.Text(), &p, &parsed)
		if err != nil {
			fmt.Println("> Err:", err)
			return
		}

		update(&p, &parsed)
	}
	printParsedLog(p)
	printScannerErr(in)
}
