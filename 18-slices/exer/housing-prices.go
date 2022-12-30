package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
	)

	var (
		locations                  []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")
	for _, el := range rows {
		data := strings.Split(el, separator)
		locations = append(locations, data[0])
		sizes = append(sizes, transformData(data[1]))
		beds = append(beds, transformData(data[2]))
		baths = append(baths, transformData(data[3]))
		prices = append(prices, transformData(data[4]))
	}

	// render Headers
	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	// render rows of Data
	for i := 0; i < len(rows); i++ {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
	}
	// render Average
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", " ", calcAvg(sizes), calcAvg(beds), calcAvg(baths), calcAvg(prices))
}

func transformData(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error during parsing: %s\n", input)
		os.Exit(1)
	}
	return num
}

func calcAvg(arr []int) float64 {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return (float64(sum)) / (float64(len(arr)))
}
