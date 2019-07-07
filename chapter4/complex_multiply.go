package chapter4

func Multiply(comp1 complex128, comp2 complex128) complex128 {
	a, b := real(comp1), imag(comp1)
	c, d := real(comp2), imag(comp2)
	abc := (a + b) * c
	cdb := (c + d) * b
	abd := (a - b) * d
	// ac - bd, ad + bc
	return complex(abc - cdb, abd + cdb)
}