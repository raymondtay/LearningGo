package main

import (
  "fmt"
  "math"
)

type Vertex struct { X, Y float64 }
type MyFloat float64 // type alias to float64

// second approach which is equivalent to the first approach
// the interesting thing about this approach is that the type MyFloat is
// associated with the Abs() method defined here. In other words, its
// decorating the type with methods....
func (f MyFloat) Abs() float64 {
  if  f < 0 {
    return float64(-f) 
  }
  return float64(f)
}

// first approach to writing the absolute function
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale (f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())

  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())

  fmt.Println(v)
  v.Scale(10)
  fmt.Println(v.Abs())
}

