package main 

import (
  "fmt"
  "math"
)

const Pi = 3.14
// The expression T(v) converts the value "v" to the type "T"

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f) 
  fmt.Println(x, y, z)
  about_constants()
  fmt.Println("1. ", needInt(Small))
  fmt.Println("2. ", needFloat(Small))
  fmt.Println("3. ", needFloat(Big))
}

const (
  // Create a huge number by shifting a 1 bit left 100 places.
  // In other words, the binary number that is 1 followed by 100 zeroes.
  Big = 1 << 100
  // Shift it right again 99 places, so we end up with 1<<1, or 2.
  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
  return x * 0.1
}

func about_constants() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")
  
  const Truth = true
  fmt.Println("Go rules?", Truth)
}


