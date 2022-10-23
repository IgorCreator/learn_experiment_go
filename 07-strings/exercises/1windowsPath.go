package main

import "fmt"

func windowsPath() {

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"

	pathLiteral := `
	c:\program files\duper super\fun.txt
	c:\program files\really\funny.png
	`

	fmt.Println(path)
	fmt.Println(pathLiteral)
}
