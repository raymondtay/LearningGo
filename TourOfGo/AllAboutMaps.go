package main

import (
  "fmt"
)
type Vertex struct {
	Lat, Long float64
}

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func function_demo() {
  pos, neg := adder() , adder()
  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-2*i))
  }
}

var m map[string]Vertex

var amap = map[string]Vertex {
  "Bell Labs" : Vertex { 1.1, 2.2},
  "Google" : Vertex {3.3, 4.4},
}
// syntax : make(map [<type of key>] <type of value>)
// "make" makes it become real
func main() {
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])

  fmt.Println(amap)

  crud_maps()

  function_demo()
}

func crud_maps() {
  m := make(map[string]int)
  m["answer"] = 1
  m["answer"] = 2
  delete(m, "answer")
  v , ok := m["answer"]
  fmt.Println("the value: ", v , "Present?", ok)
}

