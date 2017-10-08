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
		{z1, z1},
		{z1, z2},
	}

	for _, point := range cases {
		if !Eq(point.a, point.b) {
			t.Errorf("Point %v is not equal to %v", point.a, point.b)
		}
	}
}

func TestLFT(t *testing.T) {
	cases := []struct {
		a Point
		b Point
	}{
		{one, one},
		{zero, inf},
		{inf, zero},
	}

	for _, point := range cases {
		if !Eq(point.a, LFT(point.b, m2c.Matrix{0, 1, 1, 0})) {
			t.Errorf("Point %v is not the image of %v", point.a, point.b)
		}
	}
}
