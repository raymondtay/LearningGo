package main

import (
  "fmt"
  "time"
  "math"
  "math/rand"
)

// "var" declares a list of variables, as in function argument lists, the type is last.
var c, python, java bool

func add(x int, y int) int {
  return x + y
}

func swap(x, y string) (string, string) {
  return y, x 
}

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println("Welcome to the playground")
  fmt.Println("The time is", time.Now())
  fmt.Println("Random number", rand.Intn(10))
  fmt.Println(math.Pi)
  var i int = 12121212
  x := 1 // local assignment
  y := 2 // local assignment
  x = 333
  fmt.Println(add(x, y))
  a, b := swap("a", "b")
  fmt.Println("After the swap, it is:", a, b)
  fmt.Println(split(44))
  fmt.Println(i, c, python, java)
}

