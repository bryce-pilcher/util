package slice

import (
	"fmt"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	a := []interface{}{1, 2, 3, 4, 5}
	b := DeepCopy(a)
	a[1] = 6
	if b[1] == a[1] {
		t.Errorf("b := DeepCopy(%d)\na[1]=6\nb[1] == %d", a, b[1])
	} else {
		fmt.Printf("a:%d\nb:%d\n", a, b)
	}
}
