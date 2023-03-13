// go run . < ./logs/log.txt
// go run . < ./logs/log_err_missing.txt
// go run . < ./logs/log_err_negative.txt
// go run . < ./logs/log_err_str.txt

package main

import (
	"bufio"
	"os"
)

func main() {
	parsed := &result{}
	p := newParser()

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parse(in.Text(), p, parsed)
		update(p, parsed)
	}
	printParsedLog(p)
	printScannerErr([]error{in.Err(), err(p)})
}
