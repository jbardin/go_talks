package main

import "fmt"

func main() {
	// this is an int value
	x := 42

	// a pointer is declared like so
	var p *int

	// this is a pointer to x
	p = &x

	fmt.Printf("%d is stored at %p\n", x, p)
}
