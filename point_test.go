package s2

import (
	"github.com/fibo/m2c"
	"testing"
)

var z1 = Point{1 + 1i, 2}
var z2 = Point{2 + 2i, 4}
var z3 = Point{1 - 1i, 2}
var z4 = Point{2 - 2i, 4}
var inf = Infinity()
var inf2 = Point{2, 0}
var zero = Zero()
var one = Point{1, 1}

func TestConj(t *testing.T) {
	cases := []struct {
		a Point
		b Point
	}{
		{z1, z3},
		{z2, z4},
	}

	for _, point := range cases {
		if !Eq(Conj(point.a), point.b) {
			t.Errorf("Point %v is not conjugate of %v", point.a, point.b)
		}
	}
}
func TestEq(t *testing.T) {
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
		if !Eq(point.a, point.b) {
			t.Errorf("Point %v should be equal to %v", point.a, point.b)
		}
	}
}

func TestLFT(t *testing.T) {
	identity := m2c.I()
	inversion := m2c.Matrix{0, 1, 1, 0}

	cases := []struct {
		in        Point
		out       Point
		transform m2c.Matrix
	}{
		{one, one, inversion},
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

		if !Eq(point.out, got) {
			t.Errorf("Point %v is not the image of %v, result of transformation is %v", point.out, point.in, got)
		}
	}
}
