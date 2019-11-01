package main

import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello there!")
	fmt.Println(simpleGitHub.AddTwo(5, 6))

	// part 2
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = " Please give me 1 argument"
	} else {
		myString = arguments[1]
	}

	// part 3
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
