package s2

import (
	"math/cmplx"

	"github.com/fibo/m2c"
)

// A Point is represented by the couple p = (z, w) of two complex
// numbers where two points are equal if (z, w) = λ (z', w').
type Point struct {
	Z, W complex128
}

// NewPoint is a Point constructor. It only takes one Z argument, since
// W defaults to 1, i.e. the point is already normalized.
func NewPoint(Z complex128) *Point {
	return &Point{Z, 1}
}

// Conj returns the conjugated point.
func Conj(p Point) Point {
	return Point{cmplx.Conj(p.Z), cmplx.Conj(p.W)}
}

// PointEq checks that two points are equal.
func PointEq(p Point, q Point) bool {
	return p.Z*q.W == p.W*q.Z
}

// Infinity returns the point at infinity, i.e. with w = 0.
func Infinity() Point {
	return Point{1, 0}
}

// Inv implements inversion respect to a circle.
func (p *Point) Inv(c Circle) {
	var A = -c.C
	var B = complex(-c.D, 0)
	var C = complex(c.A, 0)
	var D = cmplx.Conj(c.C)

	var Z = p.Z
	var W = p.W

	p.Z = A*Z + B*W
	p.W = C*Z + D*W
}

// LFT computes a Linear fractional transformation (https://en.wikipedia.org/wiki/Linear_fractional_transformation).
func LFT(p Point, m m2c.Matrix) Point {
	return Point{m.A*p.Z + m.B*p.W, m.C*p.Z + m.D*p.W}
}

// Norm normalizes a point in the form (z, w) to the form z' or Infinity.
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
