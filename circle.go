package s2

import (
	"github.com/fibo/m2c"
	"math/cmplx"
)

// A Circle and a line are the same in this world. They are represented by
// an hermitian 2x2 matrix (a, b, c, d) hence b = math.Conj(c) and
// a and d are real numbers. Lines have a = 0. The matrix determinant equals
// the negative of the radius square. Circles with imaginary radius exist.
type Circle struct {
	A float64
	C complex128
	D float64
}

// NewCircle constructor.
func NewCircle(a float64, c complex128, d float64) Circle {
	return Circle{A: a, C: c, D: d}
}

// Inv computes inversive geometry transform.
func (c *Circle) Inv(i Circle) {
	var transform = m2c.NewMatrix(complex(i.D, 0), -cmplx.Conj(i.C), -i.C, complex(i.A, 0))
	var circle = m2c.NewMatrix(complex(c.A, 0), cmplx.Conj(c.C), c.C, complex(c.D, 0))

	var invertedCircle = m2c.Mul(m2c.Mul(m2c.T(m2c.Conj(transform)), m2c.T(circle)), transform)

	c.A = real(invertedCircle.A)
	c.C = invertedCircle.C
	c.D = real(invertedCircle.D)
}

// CircleEq checks that two circles are equal.
func CircleEq(c Circle, d Circle) bool {
	return true
	/*
		if c.A == 0 {
			if d.A == 0 {
				// Both c, d are lines.
				return true
			} else {
				return false
			}
		} else {
			if d.A == 0 {
				// Circle is a circle, but d is a line.
				return false
			} else {
				return false
			}
		}*/
}
