package main

import (
	"fmt"
)

// structures can be defined anywhere, including inside "main" or other
// functions. standard scoping rules applies to the visibilities of these
// structures.
type aStructure struct {
	person string
	height int
	weight int
}

// simple utility to create structures....notice that default arguments cannot
// be applied here.
func createAStructure(person string, height int, weight int) *aStructure {
	return &aStructure{person, height, weight}
}

func main() {
	// unnamed fields would be initialized to the default value for their
	// corresponding types
	p1 := aStructure{"fmt", 12, -2}
	p2 := aStructure{person: "ray", height: 183} //
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(*createAStructure("ray", 1, 2)) // de-referencing like in C/C++
}
