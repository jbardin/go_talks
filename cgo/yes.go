package main

import "fmt"

/*
#include <stdio.h>
int the_answer() {
	return 42;
}
*/
import "C"

func main() {
	i := C.the_answer()
	fmt.Printf("%d. Yup, that's it!\n", i)
}
