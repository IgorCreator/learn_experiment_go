package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string

	fmt.Printf("books: %T\n", books)
	fmt.Println("books: ", books)
	fmt.Printf("books: %q\n", books)
	fmt.Printf("books: %v\n", books)
	fmt.Printf("books: %#v\n", books)
	fmt.Println("--------------------")

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"
	fmt.Printf("yearly books: %#v\n", books)

	fmt.Println("--------------------")
	var (
		wBooks [winter]string
		sBooks [summer]string
	)
	wBooks[0] = books[0]
	for i := range sBooks {
		sBooks[i] = books[i+1]
	}
	fmt.Printf("winter : %#v\n", wBooks)
	fmt.Printf("summer : %#v\n", sBooks)

	fmt.Println("--------------------")
	var published [len(books)]bool
	published[0] = true
	published[len(books)-1] = true

	fmt.Println("Published Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
