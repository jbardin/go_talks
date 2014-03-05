package main

import "fmt"

func AddThree(x *int) {
	*x += 3
}

func main() {
	x := 3
	fmt.Printf("x was %d\n", x)

	AddThree(&x)
	fmt.Printf("x is now %d\n", x)
}
