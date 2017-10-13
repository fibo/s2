package s2

import (
	"github.com/fibo/m2c"
	"math/cmplx"
)

// A Point is represented by the couple p = (z, w) of two complex
// numbers where two points are equal if (z, w) = λ (z', w').
type Point struct {
	Z, W complex128
}

var conj = cmplx.Conj

// Conj returns the conjugated point.
func Conj(p Point) Point {
	q := Point{conj(p.Z), conj(p.W)}
	q.Norm()
	return q
}

// Eq checks that two points are equal.
func Eq(p Point, q Point) bool {
	p.Norm()
	q.Norm()

	return (p.Z == q.Z) && (p.W == q.W)
}

// Infinity returns the point at infinity, i.e. with w = 0.
func Infinity() Point {
	return Point{1, 0}
}

// LFT computes a Linear fraction transformation (https://en.wikipedia.org/wiki/Linear_fractional_transformation)
func LFT(p Point, m m2c.Matrix) Point {
	q := Point{m.A*p.Z + m.B*p.W, m.C*p.Z + m.D*p.W}
	return q
}

// Norm normalizes a point in the form (z, w) to the form z or Infinity.
func (p *Point) Norm() {
	if p.W == 0 {
		p.Z = 1
	} else {
		p.Z = p.Z / p.W
		p.W = 1
	}
}

// Zero returns the origin point {0, 0}.
func Zero() Point {
	return Point{0, 1}
}
