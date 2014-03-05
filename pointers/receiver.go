package main

import "fmt"

// START OMIT
type Int struct {
	X int
}

func (i Int) NotAdd(y int) { i.X += y }
func (i *Int) Add(y int)   { i.X += y }

func main() {
	i := Int{57}
	i.NotAdd(-15)
	fmt.Printf("i is still %+v\n", i)

	i.Add(-15) // but i isn't a pointer?!
	fmt.Printf("i is now %+v\n", i)
}

// END OMIT
