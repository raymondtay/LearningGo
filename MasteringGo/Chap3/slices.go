package main

import (
  "fmt"
  "sort"
)

type aStructure struct {
  person string
  age int
  weight int
}

func printAStructure() {
    mySlice := make([]aStructure, 0)
    mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
    mySlice = append(mySlice, aStructure{"Bill", 134, 45})
    mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
    mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
    mySlice = append(mySlice, aStructure{"Athina", 134, 40})
}

func main() {

    printAStructure()

    s1 := make([]int, 5)
    reSlice := s1[1:3]
    fmt.Println(s1)
    fmt.Println(reSlice)

    reSlice[0] = -100
    reSlice[1] = 123456
    fmt.Println(s1)
    fmt.Println(reSlice)

    a6 := []int{-10, 1, 2, 3, 4, 5}
    a4 := []int{-1, -2, -3, -4}
    fmt.Println("a6:", a6)
    fmt.Println("a4:", a4)
    copy(a6, a4) // copy(dest, src)
    fmt.Println("a6:", a6)
    fmt.Println("a4:", a4)
    fmt.Println()

    b6 := []int{-10, 1, 2, 3, 4, 5}
    b4 := []int{-1, -2, -3, -4}
    fmt.Println("b6:", b6)
    fmt.Println("b4:", b4)
    copy(b4, b6)
    fmt.Println("b6:", b6)
    fmt.Println("b4:", b4)

    fmt.Println()
    c6 := []int{-10, 1, 2, 3, 4, 5}
    c4 := []int{-1, -2, -3, -4}
    fmt.Println("c6:", c6)
    fmt.Println("c4:", c4)
    copy(c4, c6[0:])
    fmt.Println("c6:", c6)
    fmt.Println("c4:", c4)

}

