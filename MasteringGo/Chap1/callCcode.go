package main

//#include <stdio.h>
//void callC() {
//  printf("CAlling C Code\n");
//}
import "C"
import (
  "fmt"
)

func main() {
  fmt.Println("calling c code")
  C.callC()
  fmt.Println("Another go statement ")
}
