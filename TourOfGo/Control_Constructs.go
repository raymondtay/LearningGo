package main

import (
  "fmt"
  "math" // how to suppress compilation errors due to un-used module
  "time"
  "runtime"
)

func pow(x , n , lim float64) float64 {
  // the "if" statement can execute a statement before the condition is
  // evaluated.
  if v := math.Pow(x, n); v  < lim {
    return v
  }
  return lim
}

func runtime_info() {
  fmt.Println("Go runs on ")
  // switch without a condition is the same as saying "switch true"
  switch os := runtime.GOOS; os {
    case "darwin" : fmt.Println("OS X.")
    case "linux" : fmt.Println("Linux.")
    case "windows" : fmt.Println("Windows.")
    default: fmt.Printf("%s. \n", os)
  }
}

func decipher_time_now() {
  t := time.Now()
  switch {
  case t.Hour() < 12: fmt.Println("Good morning!")
  case t.Hour() < 17: fmt.Println("Good afternoon!")
  default: fmt.Println("Good evening!")
  }
}

// Defer: A defer statement defers the execution of a function until the
// surround function returns. The deferred call's arguments are evaluated
// immediately, but the function call is not executed until the surrounding
// function returns.
func defer_me() {
  defer fmt.Println("World!")

  fmt.Println("Hello")
}

func loop() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
    defer fmt.Println("the sum now is: ", sum) // the statements are pushed onto the stack and when this enclosing statement is completed, it would run.
  }
  fmt.Println(sum)
}

func main() {
   loop()
   fmt.Println( pow(3, 2, 10) , pow(3, 3, 20) )
   runtime_info()
   decipher_time_now()
   defer_me()
}







