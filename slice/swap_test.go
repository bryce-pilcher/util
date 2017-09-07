package slice

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	Swap(a, 1, 2)
	if a[1] != 3 {
		t.Errorf("Swap(%d, %d, %d) == %d, want %d", []interface{}{1, 2, 3, 4}, 1, 2, a, []interface{}{1, 3, 2, 4})
	} else {
		fmt.Printf("TestSwap%d\n", a)
	}
}
