package rand

import (
	"math/rand"
	"time"
)

/*
   This function returns a random integer in range (a,b).
   It expects that a <= b, however it will work if the order is reversed.
*/
func RangeInt(a int, b int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	if a > b {
		return int(rand.Float64()*float64(b-a)) + a
	} else {
		return int(rand.Float64()*float64(a-b)) + b
	}
}
