package s2

import (
	"github.com/fibo/m2c"
	"math/cmplx"
)

var abs = cmplx.Abs
var conj = cmplx.Conj

// A Circle and a line are the same in this world.
// They are represented by an hermitian 2x2 matrix (a, b, c, d)
// hence b = math.Conj(c) and a and d are real numbers.
// Lines have a = 0.
// The matrix determinant equals the negative of the radius square.
// When determinant is zero, the circle degenerates to a line.
// Imaginary circles exist, they have negative radius.
type Circle struct {
	A float64
	C complex128
	D float64
}

// NewCircle constructor. When radius is zero it will degenerate to a
// point, when radius is negative it will create an imaginary circle.
func NewCircle(center complex128, radius float64) Circle {
	// TODO check this matrix
	return Circle{
		A: 1,
		C: -center,
		D: abs(center) - (radius * radius),
	}
}

// NewLine constructor, returns a line that is the locus of points
// equidistant from the given p, q points.
func NewLine(p complex128, q complex128) Circle {
	// TODO check this matrix
	return Circle{
		A: 0,
		C: p - q,
		D: 1,
	}
}

// Inv computes inversive geometry transform.
func (c *Circle) Inv(i Circle) {
	var transform = m2c.NewMatrix(complex(i.D, 0), -conj(i.C), -i.C, complex(i.A, 0))
	var circle = m2c.NewMatrix(complex(c.A, 0), conj(c.C), c.C, complex(c.D, 0))

	var invertedCircle = m2c.Mul(m2c.Mul(m2c.T(m2c.Conj(transform)), m2c.T(circle)), transform)
	c.A = real(invertedCircle.A)
	c.C = invertedCircle.C
	c.D = real(invertedCircle.D)
}

// CircleEq checks that two circles are equal.
func CircleEq(c Circle, d Circle) bool {
	if c.A == 0 && d.A == 0 {
		// Both c, d are lines.
		return true
	}

	return false
	// TODO
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
