package main

import "fmt"

func main() {
	const (
		ETH = 9 - iota // 9
		WAN            // 8
		ICX            // 7
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
		ICX: 20,
	}

	// uses well-defined names (ETH, WAN, ...) - good
	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
	fmt.Printf("1 BTC is %g ICX\n", rates[ICX])

	fmt.Printf("%#v\n", rates)
}
