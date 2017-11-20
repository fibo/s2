package s2

import (
	"testing"
)

func TestInv(t *testing.T) {
	var unit1 = NewCircle(0, 1)
	var unit2 = NewCircle(0, 1)
	var unit3 = NewCircle(0, 1)

	cases := []struct {
		in        Circle
		out       Circle
		inversion Circle
	}{
		{unit1, unit2, unit3},
	}

	for _, circle := range cases {
		circle.in.Inv(circle.inversion)

		if !CircleEq(circle.out, circle.in) {
			t.Errorf("Circle %v is not the inversion of %v", circle.out, circle.in)
		}
	}
}
