package main

type T struct{}

func (t *T) Foo() {}

type Fooer interface {
	Foo()
}

func main() {
	t := T{}
	f := Fooer(t)
	f.Foo()
}
