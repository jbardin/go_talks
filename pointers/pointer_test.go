package main

import "testing"

type BigStruct struct {
	padding [96]byte
	X, Y    int
}

// These are for nothing more than doing a small read/write.
func (p BigStruct) ValueAdd() {
	p.X += p.Y
}

func (p *BigStruct) PointerAdd() {
	p.X += p.Y
}

func BunchOfBigStructs(n int) []BigStruct {
	var s []BigStruct
	for i := 0; i < n; i++ {
		p := BigStruct{
			X: i,
			Y: i + 1,
		}
		s = append(s, p)
	}
	return s
}

type LittleStruct struct {
	W int
	X int
	Y int
	Z int
}

func (p LittleStruct) ValueAdd() {
	p.X += p.Y
}

func (p *LittleStruct) PointerAdd() {
	p.X += p.Y
}

func BunchOfLittleStructs(n int) []LittleStruct {
	var s []LittleStruct
	for i := 0; i < n; i++ {
		p := LittleStruct{
			X: i,
			Y: i + 1,
		}
		s = append(s, p)
	}
	return s
}

func BenchmarkLittle__Value(b *testing.B) {
	points := BunchOfLittleStructs(100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(points); j++ {
			points[j].ValueAdd()
		}
	}
}

func BenchmarkLittlePointer(b *testing.B) {
	points := BunchOfLittleStructs(100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(points); j++ {
			points[j].PointerAdd()
		}
	}
}
func BenchmarkBig__Value(b *testing.B) {
	points := BunchOfBigStructs(100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(points); j++ {
			points[j].ValueAdd()
		}
	}
}

func BenchmarkBigPointer(b *testing.B) {
	points := BunchOfBigStructs(100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(points); j++ {
			points[j].PointerAdd()
		}
	}
}
