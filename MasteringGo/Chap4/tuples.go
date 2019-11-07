package main

import (
	"fmt"
)

func retThree(x int) (int, int, int) {
	return x, x + 1, x + 2
}

func main() {
	a, b, c := retThree(4) // unpack the tuples otherwise compilation error
	fmt.Println(">>", a, b, c)
}
