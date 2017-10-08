package s2

// Circles and lines are the same in this world. They are represented by
// an hermitian 2x2 matrix (a, b, c, d) hence b = math.Conj(c) and
// a and d are real numbers. Lines has a = 0.
type Circle struct {
	A float64
	C complex128
	D float64
}

// TODO circle.Inv() and point.Inv() computes inversive geometry transform.
