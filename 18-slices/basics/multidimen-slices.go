package main

import (
	"fmt"
	"strconv"
	"strings"
)

const spendingData = `200 100
		25 10 45 60
		5 15 35
		95 10
		50 25`

func main() {
	spendings := convertDataInSlice(spendingData)

	for i, daily := range spendings {
		var total int
		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func convertDataInSlice(content string) [][]int {
	lines := strings.Split(content, "\n") // splits the string by new lines

	spendings := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line) // splits the string by white space

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			spending, _ := strconv.Atoi(field) // convert string in int
			spendings[i][j] = spending
		}
	}

	return spendings
}
