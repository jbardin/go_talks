package main

import (
	"log"
	"unsafe"
)

func fubar(x *int) *int {
	i := uintptr(unsafe.Pointer(x)) + 16
	return (*int)(unsafe.Pointer(i))
}

func main() {
	s := []int{1, 2, 3}
	log.Println(*(fubar(&s[0])))
}
