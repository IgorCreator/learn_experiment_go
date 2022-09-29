package main

import "fmt"

func trigger_sound() {
	// only this file can access the imported fmt package
	// when others also do so, they can also access
	//   their own `fmt` "name"

	fmt.Println("Wow!!! you in Library Package")
	methodInExecutablePackage()
}
