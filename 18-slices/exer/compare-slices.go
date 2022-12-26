package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(namesA)
	sort.Strings(namesB)

	fmt.Printf("%q\n", namesA)
	fmt.Printf("%q\n", namesB)
	fmt.Printf("Slices DeepEqual: %v\n", reflect.DeepEqual(namesA, namesB))
	compareSlices(namesA, namesB)
}

func compareSlices(namesC []string, namesB []string) {
	if len(namesC) == len(namesB) {
		for i := range namesC {
			if namesC[i] != namesB[i] {
				return
			}
		}
		fmt.Println("They are equal.")
	}
}
