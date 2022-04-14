/*
Methods

In fact, you can define a method on any type you define in your package, not just structs.

You cannot define a method on a type from another package, or on a basic type.
*/

package main

import (
	"fmt"
	"math"
)
type myfloat float64
func (f myfloat) Abs () float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f:=myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
