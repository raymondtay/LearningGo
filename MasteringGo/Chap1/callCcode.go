package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// To build the entire application, perform the following:
// (1) mkdir callClib
// (2) mv all *.h and *.c into callClib
// (3) run `gcc -c callClib/*.c`
// (4) run `ar rs callC.a *.o`
// (5) go run callCcode.go
func main() {
	fmt.Println("calling c code")
	C.cHello()
	fmt.Println("Another go statement ")
	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Ray!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")
}
