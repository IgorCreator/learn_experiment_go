package main

import "fmt"

func main() {
	names := [3]string{"Einstein", "Tesla", "Shepard"}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	ratios := [1]float64{3.14145}
	alives := [...]bool{true, false, true, false}
	zero := [0]byte{}

	fmt.Println("ARRAYS:")
	fmt.Printf("names    : %[1]T %[1]q\n", names)
	fmt.Printf("distances: %[1]T %[1]d\n", distances)
	fmt.Printf("data     : %[1]T %[1]d\n", data)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero)

	fmt.Println("\nSLICES:")
	namesArr := []string{"Einstein", "Tesla", "Shepard"}
	distancesArr := []int{50, 40, 75, 30, 125}
	dataArr := []byte{'H', 'E', 'L', 'L', 'O'}
	ratiosArr := []float64{3.14145}
	alivesArr := []bool{true, false, true, false}
	zeroArr := []byte{}
	fmt.Printf("names    : %T %d %t\n", namesArr, len(namesArr), namesArr == nil)
	fmt.Printf("distances: %T %d %t\n", distancesArr, len(distancesArr), distancesArr == nil)
	fmt.Printf("data     : %T %d %t\n", dataArr, len(dataArr), dataArr == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratiosArr, len(ratiosArr), ratiosArr == nil)
	fmt.Printf("alives   : %T %d %t\n", alivesArr, len(alivesArr), alivesArr == nil)
	fmt.Printf("zero     : %T %d %t\n", zeroArr, len(zeroArr), zeroArr == nil)
}
