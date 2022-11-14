package main

import "fmt"

//  go run main.go

func main() {
	firstLoopSteps()

	fmt.Println("---------------")
	whileLoopWithBreak()

	fmt.Println("---------------")
	whileLoopWithContinue()

	fmt.Println("---------------")
	nestedLoops()
}

func nestedLoops() {
	//var (
	//	n = 5
	//	m = 4
	//)
	const max = 5

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}

}

func whileLoopWithContinue() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			// just by putting this here we solve the problem
			i++
			continue
		}

		sum += i

		fmt.Println(i, "-->", sum)

		i++
	}

	fmt.Println(sum)
}

func whileLoopWithBreak() {
	var (
		sum int
		i   = 1
	)

	for {
		if sum >= 10 {
			break
		}
		sum += i
		fmt.Println(i, " --> ", sum)
		i++
	}
}

func firstLoopSteps() {
	var sum int

	for i := 1; i <= 5; i++ {
		sum += 2
		fmt.Println(sum)
	}
}
