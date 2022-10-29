package main

import (
	"fmt"
	"os"
)

//  go build -o args-app 03-arg-count.go
// 	./args-app
// 	./args-app hi
// 	./args-app hi 10
// 	./args-app more args int app

func main() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf(
			`There are two: "%s %s"`+"\n",
			args[1], args[2],
		)
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}
