package main
import (
"fmt"
)
type Digit int
type Power2 int

const PI = 3.1415926

// the syntax is weird...with rules and exceptions.
// sad...
const (
c1 = "c1c1c1"
c2 = "c2c2c2"
c3 = "c3c3c3"

p2_0 Power2 = 1 << iota
_
p2_2
_
p2_4

)


func main() {
  fmt.Println(">> ", PI)
  fmt.Println(">> ", c1)
  fmt.Println(">> ", c2)
  fmt.Println(">> ", c3)

  fmt.Println("2^0: " , p2_0)
  fmt.Println("2^2: " , p2_2)
  fmt.Println("2^4: " , p2_4)
}

