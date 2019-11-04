package main

import (
	"fmt"
)

func d1() {
	for i := 4; i > 0; i-- {
		defer fmt.Print(i, "+")
	}
}

// deferred functions are always evaluated after their surrounding function
// terminates. In this case, its the for-loop.
func d2() {
	for i := 4; i > 0; i-- {
		defer func() { // anonymouns function that consumes no value.
			fmt.Print(i, "*")
		}()
	}
	fmt.Println()
}

// deferred functions are always evaluated after their surrounding function
// terminates. In this case, its the for-loop and the main difference is that
// the value for each iteration is already bound to the deferred function which
// implies that there's an least 1 copy of the function and as many as n-values
// or i-values are stored somewhere to be executed.
func d3() {
	for i := 4; i > 0; i-- {
		defer func(n int) { // anonymouns function that consumes a value in each iteration
			fmt.Print(n, "-")
		}(i)
	}
	fmt.Println()
}

func main() {
	d1()
	d2()
	fmt.Println("<<")
	d3()
	fmt.Println("<<")
}
