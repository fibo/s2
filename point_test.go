package s2

import (
	"testing"
)

var z1 = Point{1, 2}
var z2 = Point{2, 4}
var inf = Infinity()

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
