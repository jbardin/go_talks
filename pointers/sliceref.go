package main

import "fmt"

// START_1 OMIT
func badAppend(s []int, val int) {
	s = append(s, val)
}

// END_1 OMIT

func main() {
	sliceOfInt := make([]int, 3, 4)
	sliceOfInt[0] = 0
	sliceOfInt[1] = 1
	sliceOfInt[2] = 2

	// START_2 OMIT
	fmt.Println("sliceOfInt:", sliceOfInt)

	badAppend(sliceOfInt, 3)
	fmt.Println("It's there!:", sliceOfInt[:4])

	// Did that really work?
	// fmt.Println("No, not really:", sliceOfInt)
	// END_2 OMIT
}
