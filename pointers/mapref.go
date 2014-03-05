package main

import "fmt"

// START OMIT
func PointlessInsert(m map[string]string, key, val string) {
	m[key] = val
}

func main() {
	m := make(map[string]string)
	fmt.Printf("Map is empty, %+v\n", m)

	PointlessInsert(m, "apple", "juice")
	fmt.Printf("Map is juicy, %+v\n", m)
}

// END OMIT
