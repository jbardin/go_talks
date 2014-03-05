package main

import "fmt"

// Pointless insert function
func insert(m map[string]string, key, val string) {
	m[key] = val
}

func main() {
	m := make(map[string]string)
	fmt.Printf("Map is empty, %+v\n", m)

	insert(m, "apple", "juice")
	fmt.Printf("Map is juicy, %+v", m)
}
