package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go run ./exer/log-analyzer.go < ./exer/logs/log.txt
//go run ./exer/log-analyzer.go < ./exer/logs/log_err_missing.txt
//go run ./exer/log-analyzer.go < ./exer/logs/log_err_negative.txt
//go run ./exer/log-analyzer.go < ./exer/logs/log_err_str.txt

func main() {
	in := bufio.NewScanner(os.Stdin)

	var sum, qtyLines int
	logs := make(map[string]int)
	for in.Scan() {
		qtyLines++
		line := strings.Fields(in.Text())
		if len(line) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", line, qtyLines)
			return
		}

		domain := strings.ToLower(line[0])
		visit, err := strconv.Atoi(line[1])
		if visit < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", line[1], qtyLines)
			return
		}
		sum += visit

		if _, existsInMap := logs[domain]; existsInMap {
			visit += logs[domain]
		}

		logs[domain] = visit
	}

	fmt.Println(fmt.Sprintf("%-25s %10s", "DOMAIN", "VISIT"))
	fmt.Println(strings.Repeat("-", 40))
	for domain, visits := range logs {
		fmt.Printf("%-25s %10d\n", domain, visits)
	}

	fmt.Println(fmt.Sprintf("\n%-25s %10d", "Total", sum))

	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
