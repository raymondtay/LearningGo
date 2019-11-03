package main

import (
  "fmt"
  "unsafe"
  "runtime"
  "time"
)

func printStats(mem runtime.MemStats) {
  runtime.ReadMemStats(&mem)
  fmt.Println("mem.Alloc:", mem.Alloc)
  fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
  fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
  fmt.Println("mem.NumGC:", mem.NumGC)
  fmt.Println("-------")
}


func demoGC() { 
  var mem runtime.MemStats
  printStats(mem)
  for i := 0; i < 10;  i++ {
    s := make([]byte, 500000)
    if s == nil {
      fmt.Println("OPeration failed!!!")
    }
    time.Sleep(3 * time.Second)
  }
  printStats(mem)
}

func  main() {
  var value int64 = 5
  var p1 = &value
  var p2 = (*int32) (unsafe.Pointer(p1)) // create a int32 pointer that points to a int64 value
  fmt.Println("*p1: ", *p1)
  fmt.Println("*p2: ", *p2)
  *p1 = 132313313
  fmt.Println("*p1: ", *p1)
  fmt.Println("*p2: ", *p2)

  // demoGC() // demonstration of the Garbage Collection in Go

  moreUnsafe() // more unsafe stuff !
}

func moreUnsafe() {
  array := [...]int{0,1,-2,3,4}
  pointer := &array[0]
  fmt.Print(*pointer, " ")
  memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

  for i := 0; i < len(array) -1; i++ {
    pointer = (*int)(unsafe.Pointer(memoryAddress))
    fmt.Print(*pointer, " ")
    memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
  }

  fmt.Println()
  pointer = (*int)(unsafe.Pointer(memoryAddress))
  fmt.Print("One more: " , *pointer, " ")
  memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
  fmt.Println()
}
