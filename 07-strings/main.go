package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	//Example 1
	eq := "1 + 2 = "
	sum := 1 + 2
	// string concat op can only be used with strings
	// you need to convert it using strconv.Itoa
	// Itoa = Integer to ASCII
	fmt.Println(eq + strconv.Itoa(sum))

	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)

	//Example 2
	name, last := "carl", "sagan"
	name += " edward"
	fmt.Println(name + " " + last)

	//Example 3
	rawStrLiteral()

	//Example 4
	msg := "Hello"
	l := len(msg)
	s := strings.ToUpper(msg) + strings.Repeat("!", l)
	outputString(s, len(s), utf8.RuneCountInString(s))

	f := "ƒ(ß+∆)=ƒ(∂)"
	outputString(f, len(f), utf8.RuneCountInString(f))

	pi := "π=3.14"
	outputString(pi, len(pi), utf8.RuneCountInString(pi))

	//Example 4
}

func rawStrLiteral() {
	litStr := "\t<html>\n\t\t<body>\n\t\t\t<h1>Hello!</h1>\n\t\t</body>\n\t</html>"
	rawStr := `
	<html>
		<body>
			<h1>Hello!</h1>
		</body>
	</html>
	`
	fmt.Println(litStr)
	fmt.Println(rawStr)
}

func outputString(str string, bytesLen int, symLen int) {
	fmt.Println("----------")
	fmt.Println("String: " + str)
	fmt.Println("Bytes len: ", bytesLen)
	fmt.Println("Symbols len: ", symLen)
}
