package main

import (
	"fmt"
)

func main() {

	var iMap map[string]int
	var anotherMap map[string]int

	// one way to create a empty map where (k,v) = (string, int)
	iMap = make(map[string]int)

	// another way to create a map
	anotherMap = map[string]int{
		"k1": 1,
		"k2": 2, // STUPID syntax !!!!
	}
	for key, value := range iMap {
		fmt.Println(key, value)
	}
	for key, value := range anotherMap {
		fmt.Println(key, value)
	}
	fmt.Println(">>", iMap)
	fmt.Println(">>", anotherMap)
}
