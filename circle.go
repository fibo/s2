package s2

// Circles and lines are the same in this world. They are represented by
// an hermitian 2x2 matrix (a, b, c, d) hence b = math.Conj(c) and
// a and d are real numbers
type Circle struct {
	a float64
	c complex128
	d float64
}

// TODO circle.Inv() and point.Inv() computes inversive geometry transform.
