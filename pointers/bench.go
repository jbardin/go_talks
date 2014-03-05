package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "test", "-bench", ".", "pointer_test.go")
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

// copied from the test we're running for display in the slide

// START OMIT
type LittleStruct struct {
	W, X, Y, Z int
}

func (p LittleStruct) ValueAdd()    { p.X += p.Y }
func (p *LittleStruct) PointerAdd() { p.X += p.Y }

type BigStruct struct {
	padding [96]byte
	X, Y    int
}

func (p BigStruct) ValueAdd()    { p.X += p.Y }
func (p *BigStruct) PointerAdd() { p.X += p.Y }

// END OMIT
