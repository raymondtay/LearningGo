package main

import "fmt"

type Vertex struct {
    X int
    Y int
  }

// The experience of the pointers resembles a lot like the C-programming
// language.
func main() {
  i , j :=  42, 2701

  p := & i
  fmt.Println(*p)

  *p = 21
  fmt.Println(i)

  p = &j
  *p = *p / 37 // re-assignment of pointer after an divide-operation.
  fmt.Println(j)

  v, y := 1, 3
  vertice := Vertex{v, y}
  pv := &vertice // pointer to the vertice


  fmt.Println(Vertex{1, 2}) // printing a "Vertex"
  fmt.Println(*pv) // de-referencing the pointer
  pv.X = 3448
  fmt.Println(*pv) // de-referencing the pointer

  arrays_demo()
  array_maker()
  range_demo()
}

type IPAddr [4]byte
func (p IPAddr) String() string {
  return fmt.Sprintf("%")
}

func arrays_demo() {
  var a [2]string // 2-element string array
  a[0] = "hello"
  a[1] = "world"
  fmt.Println(a[0],a[1])
  fmt.Println(a)
  primes := []int{2,3,5,7,11,13} // omitting the size of the arrays seems to be quite interesting.
  subprimes := primes[1:4]
  fmt.Println(primes, subprimes)

  s := []struct {
    X int
    Y int
    F bool
  } {
    {1, 2, false},
    {3, 2, false},
    {3, 4, false},
    {6, 2, false},
    {5, 5, false},
    {2, 3, true }, // last element needs to have a comma, ... clearly a compiler problem
  }
  fmt.Println(s)
  fmt.Println(s[:2])
  fmt.Println(s[2:])
}

func array_maker() {
  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func range_demo() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i) // == 2**i
  }
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}

