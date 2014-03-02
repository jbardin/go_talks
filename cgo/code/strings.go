package main

import (
	"fmt"
	"unsafe"
)

// START OMIT

/*
// careful, this might cause a buffer overrun
char charArray[16] = "sixteen chars xx";
char a[16]         = "!!BUFFER OVERRUN";
char* myString = (char*)&charArray;
#include <string.h>
*/
import "C"

func main() {
	// To get the string into GoString, we need a *char
	// but charArray isn't null terminated, so don't do this!
	myString := C.GoString((*C.char)(unsafe.Pointer(&C.charArray)))
	fmt.Println(myString)

	// Assigning to a char array requires constant sized arrays
	C.charArray = *(*[16]C.char)(unsafe.Pointer(&[16]byte{'f', 'r', 'o', 'm', ' ', 'g', 'o'}))
	fmt.Println(C.GoString((*C.char)(unsafe.Pointer(&C.charArray))))

	// Using strncpy still needs more unsafe pointers type assertions
	C.strncpy((*C.char)(unsafe.Pointer(&C.charArray)), C.CString("my string"), 9)
	fmt.Println(C.GoString((*C.char)(unsafe.Pointer(&C.charArray))))
}

// END OMIT
