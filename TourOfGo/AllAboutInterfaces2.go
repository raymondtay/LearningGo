package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
}

type T struct {
  S string
  X int
  Y uint 
}

// This methods means type T implements the interface I, but we don't need to
// explicitly declare that it does so.
func (t T) M() {
  fmt.Println(t.S)
  //fmt.Println(t.X)
  //fmt.Println(t.Y)
}
func main() {
  // the expression or statement `var i I` can be replaced by `var i` given
  // that `T{...}` is enough for the compiler to understand what it is.
  // var i I = T {"hello world"}
  // i.M()

  var i I 
  i =  &T {"hello"}
  describe(i)
  i.M()

  i = F(math.Pi)
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Println("(%v, %T)\n", i, i)
}

func (t *T) M() { // this is considered a re-declaration of the method.
  fmt.Println(t.S) 
}

type F float64

func (f F) M() {
  fmt.Println(f)
}

