package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
  Name string
  Age int
}

// "String()" is maciam like "toString()" in Java/Scala
func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  var p Person
  p = Person{"ray", 42}
  fmt.Println(p.String())
  fmt.Println(p)
  do(21)
  do("hello")
  do(true)
}


