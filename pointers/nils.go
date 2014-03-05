package main

import "fmt"

type T struct {
	x int
}

func (t *T) Foo() {
	t.x++
}

// START OMIT
type Fooer interface {
	Foo()
}

func main() {
	// t is declared, but it's nil
	var t *T      // t is nil
	f := Fooer(t) // t is a Fooer, so we can convert it

	// later on we want to see if we have a t
	if f == nil {
		fmt.Println("OMG, no T! ABORT!")
		return
	}

	fmt.Println("It's ok, we have a Fooer,")
	f.Foo()
}

// END OMIT
