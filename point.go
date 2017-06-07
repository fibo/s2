// The Riemann Sphere (https://en.wikipedia.org/wiki/Riemann_sphere)
// extends the complex plane by adding one point at Infinity.
// Actually is the plane plus one point, say it the North pole,
// with some good math formulas to ensure nice behaviuors, in
// particular to avoid functions discontinuity and other goodies like
// conformal map (https://en.wikipedia.org/wiki/Conformal_map) i.e.
// circles are mapped into circles.
// I dedicate this work to my father Umberto Casati and the great
// mathematician Luigi Bianchi who was a great inspiration for me.
package s2

// Every point is represented by the couple p = (z, w) of two complex
// numbers where two points are equal if (z, w) = λ (z', w').
type Point struct {
	Z, W complex128
}

// Infinity returns the point at infinity, i.e. with w = 0.
func Infinity() Point {
	return Point{1, 0}
}

// Eq checks that two points are equal.
func Eq(p Point, q Point) bool {
	p.Norm()
	q.Norm()

	return (p.Z == q.Z) && (p.W == q.W)
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
