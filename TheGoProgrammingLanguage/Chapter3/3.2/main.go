package main

import (
	"fmt"
	"math"
)

// use of float32 can result in errors
// the most recomended type is float64

// Floating-point values are conveniently printed with Printf’s %g verb,
// which chooses the most compact representation that has adequate
// precision, but for tables of data, the %e (exponent) or %f (no exponent)
// forms may be more appropriate
func main() {
	var f float32 = 16777216
	fmt.Println(f == f+1) // true

	for x := 0; x < 8; x++ {
		// 3 decimal digits, aligned in 8 character field
		fmt.Printf("x = %d e = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //  "0 -0 +Inf -Inf NaN"
}
