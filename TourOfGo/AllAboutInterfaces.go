package main

import (
  "fmt"
  "math"
)


type Vertex struct { X , Y float64 }
type MyFloat float64


// An interface type is defined as a set of method signautres.
// A value of interface type can hold any value that implements those methods.
type Abser interface {
  Abs() float64
}

func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3,4}

  a = f // a MyFloat implements Abser
  fmt.Println(a.Abs())
  a = &v // a *Vertex implements Abser

  // v is a Vertex (not a *Vertex) and does not implement Abser
  a = v
  fmt.Println(a)
  fmt.Println(a.Abs())
}

func (f MyFloat) Abs() float64 {
  if  f < 0 {
    return float64(-f) 
  }
  return float64(f)
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

