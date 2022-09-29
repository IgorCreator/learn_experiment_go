# Cheatsheet: Write your First Go Program

---

## COMMAND LINE COMMANDS:

* Enter into a directory: `cd directoryPath`

* **WINDOWS:**

    * List files in a directory: `dir`

* **OS X & LINUXes:**

    * List files in a directory: `ls`

## BUILDING & RUNNING GO PROGRAMS:

* **Build a Go program:**

    * While inside a program directory, type:
        * `go build main.go`

* **Run a Go program:**

    * While inside a program directory, type:
        * `go run main.go`

## WHERE YOU SHOULD PUT YOUR SOURCE FILES?

* In any directory you like.

## FIRST PROGRAM

### Create a directory
* Create a new directory:
    * `mkdir myDirectoryName`
* Go to that directory:
    * `cd myDirectoryName`

### Add the source code files
* Create a new `code main.go` file.
    * This is going to the create the file in the folder, and open up it in the Visual Studio Code.
* And add the following code to it and save it.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

### Run the program
* Finally, return back to the command-line.
    * Run it like this: `go run main.go`
* If you create other files and run them all, you can use this command:
    * `go run .`

---
  1. Make your first app: **Say hello to yourself.**

  2. Build your program using `go build`

  3. **Send it to your friend**

     S/he should use be using the same operating system.

     For example, if you're using Windows, then hers/his should be Windows as well.

  4. **Send your program to a friend with a different operating system**

     So, you should compile your program for her operating system.

     **Create an OSX executable:**
     `GOOS=darwin GOARCH=386 go build`

     **Create a Windows executable:**
     `GOOS=windows GOARCH=386 go build`

     **Create a Linux executable:**
     `GOOS=linux GOARCH=arm GOARM=7 go build`

     **You can find the full list in here:**
     https://golang.org/doc/install/source#environment

     **NOTE:** If you're using the command prompt or the PowerShell, you may need to use it like this:
     `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

  5. **Call [Print](https://golang.org/pkg/fmt/#Print) instead of [Println](https://golang.org/pkg/fmt/#Println)** to see what happens.

  6. **Call [Println](https://golang.org/pkg/fmt/#Println) or [Print](https://golang.org/pkg/fmt/#Print) with multiple values** by separating them using commas.

  7. **Remove the double quotes from a string literal** and see what happens.

  8. **Move the package and import statement** to the bottom of the file and see what happens.

  9. **[Read Go online documentation](https://golang.org/pkg)**.

  10. Take a quick look at the packages and read what they do.

  11. Look at their source-code by clicking on their titles.

  13. Also, **take a tour on**: https://tour.golang.org/

  14. Have a quick look. Check out the language features.