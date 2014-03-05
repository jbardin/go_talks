// If x is a pointer to a struct, x.y is shorthand for (*x).y

type T struct {
	X int
}

func (t T) M1()  {}
func (t *T) M2() {}

tPointer := &T{1}
tValue := T{2}

tPointer.X    // instead of (*tPointer).X
tPointer.M1() // instead of (*tPointer).M1() OR T.M1(*tPointer) 
tValue.X      // no magic here
tValue.M2()   // instead of (&tValue).M2() OR (*T).M1(&tValue)
