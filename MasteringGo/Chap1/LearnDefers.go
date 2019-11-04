package main

import (
	"fmt"
	"runtime"
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
func systemInfo() {
	fmt.Println("------->")
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("------->")
}
func main() {
	systemInfo()
	d1()
	d2()
	fmt.Println("<<")
	d3()
	fmt.Println("<<")
}
