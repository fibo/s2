package s2

// A Circle and a line are the same in this world. They are represented by
// an hermitian 2x2 matrix (a, b, c, d) hence b = math.Conj(c) and
// a and d are real numbers. Lines has a = 0.
type Circle struct {
	A float64
	C complex128
	D float64
}

// NewCircle constructor.
func NewCircle(a float64, c complex128, d float64) Circle {
	return Circle{A: a, C: c, D: d}
}

// TODO circle.Inv() computes inversive geometry transform.
