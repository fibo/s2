package s2

import (
	"github.com/fibo/m2c"
	"testing"
)

func TestPointConj(t *testing.T) {
	var z1 = Point{1 + 1i, 2}
	var z2 = Point{2 + 2i, 4}
	var z3 = Point{1 - 1i, 2}
	var z4 = Point{2 - 2i, 4}

	cases := []struct {
		a Point
		b Point
	}{
		{z1, z3},
		{z2, z4},
	}

	for _, point := range cases {
		if !PointEq(Conj(point.a), point.b) {
			t.Errorf("Point %v is not conjugate of %v", point.a, point.b)
		}
	}
}
func TestPointEq(t *testing.T) {
	var inf = Infinity()
	var inf2 = Point{2, 0}
	var z1 = Point{1 + 1i, 2}
	var z2 = Point{2 + 2i, 4}

	cases := []struct {
		a Point
		b Point
	}{
		{inf, inf},
		{inf, inf2},
		{z1, z1},
		{z1, z2},
	}

	for _, point := range cases {
		if !PointEq(point.a, point.b) {
			t.Errorf("Point %v should be equal to %v", point.a, point.b)
		}
	}
}

func TestPointLFT(t *testing.T) {
	var identity = m2c.I()
	var inversion = m2c.Matrix{A: 0, B: 1, C: 1, D: 0}
	var inf = Infinity()
	var inf2 = Point{2, 0}
	var z1 = *NewPoint(2 + 8i)
	var zero = Zero()
	var one = NewPoint(1)

	cases := []struct {
		in        Point
		out       Point
		transform m2c.Matrix
	}{
		{*one, *one, inversion},
		{zero, inf, inversion},
		{inf, zero, inversion},
		{zero, inf2, inversion},
		{inf2, zero, inversion},
		{z1, z1, identity},
		{inf, inf, identity},
		{inf, inf2, identity},
	}

	for _, point := range cases {
		got := LFT(point.in, point.transform)

		if !PointEq(point.out, got) {
			t.Errorf("Point %v is not the image of %v, result of transformation is %v", point.out, point.in, got)
		}
	}
}

func TestPointInv(t *testing.T) {
	var unit = NewCircle(1, 0, 1)
	var z1 = NewPoint(0)
	var inf = Infinity()
	var z2 = NewPoint(0)

	cases := []struct {
		in        Point
		out       Point
		inversion Circle
	}{
		{*z1, inf, unit},
		{inf, *z2, unit},
	}

	for _, point := range cases {
		point.in.Inv(point.inversion)

		if !PointEq(point.out, point.in) {
			t.Errorf("Point %v is not the inversion of %v", point.out, point.in)
		}
	}
}
