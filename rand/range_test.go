package rand

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	cases := []struct {
		a, b int
	}{
		{1, 9},
		{1, 2},
		{100, 1000},
	}
	for _, c := range cases {
		x := RangeInt(c.a, c.b)

		if x < c.a || c.b < x {
			t.Errorf("RangeInt(%d,%d) == %d, want %d <= x <= %d\n", c.a, c.b, x, c.a, c.b)
		} else {
			fmt.Printf("%d <= RangeInt(%d, %d) == %d <= %d\n", c.a, c.a, c.b, x, c.b)
		}
	}
}
