package main

import (
	"fmt"
	"unsafe"
)

// START OMIT

/*
// careful, this might cause a buffer overrun
char charArray[17] = "this is my string" ;
#include <string.h>
*/
import "C"

func main() {
	// To get the string into GoString, we need a *char
	// but this one might not be null-terminated
	myString := C.GoString((*C.char)(unsafe.Pointer(&C.charArray)))
	fmt.Println(myString, "and it didn't cause a crash")

	// Assigning to a char array requires constant sized arrays
	C.charArray = *(*[17]C.char)(unsafe.Pointer(&[17]byte{'f', 'r', 'o', 'm', ' ', 'g', 'o'}))
	fmt.Println(C.GoString((*C.char)(unsafe.Pointer(&C.charArray))))

	// Using strncpy still needs more unsafe pointers type assertions
	C.strncpy((*C.char)(unsafe.Pointer(&C.charArray)), C.CString("my string"), 9)
	fmt.Println(C.GoString((*C.char)(unsafe.Pointer(&C.charArray))))
}

// END OMIT
