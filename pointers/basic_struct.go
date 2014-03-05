type T struct {
	X int
}

// this is a literal T
t := T{X: 2}

// this is a pointer to a T
tp := &T{X: 3}

// since tp is already of type *T...
tp = &t

// incidentally, these are basically equivalent
tp = &T{}
tp = new(T)
