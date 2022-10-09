package main

import "fmt"

func printDecNumbersFromHex() {
	//10 in hexadecimal
	fmt.Println(0xa)

	// 16 in hexadecimal
	// 0x10
	//   ^^-----  1 * 0 = 0
	//   |
	//   +------ 16 * 1 = 16
	//                  = 16
	fmt.Println(0x10)

	// I'm going to print 150 in hexadecimal
	// 0x96
	//   ^^-----  1 * 6 = 6
	//   |
	//   +------ 16 * 9 = 144
	//                  = 150
	fmt.Println(0x96)

	//  1. Print 0 to 9 in hexadecimal
	fmt.Println(0x1, 0x2, 0x9)
	//  2. Print 15 in hexadecimal
	fmt.Println(0xf)
	//  3. Print 17 in hexadecimal
	fmt.Println(0x11)
	//  4. Print 25 in hexadecimal
	fmt.Println(0x19)
	//  5. Print 50 in hexadecimal
	fmt.Println(0x32)
	//  6. Print 100 in hexadecimal
	fmt.Println(0x64)
}
