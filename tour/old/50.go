/*
Methods

Go does not have classes. However, you can define methods on struct types.

The method receiver appears in its own argument list between the func keyword and the method name.
*/

package main

import (
	"fmt"
	"math"
)

type vec struct {
	x,y float64
}


func (v *vec) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v:=&vec{3,4}
	fmt.Println(v.Abs())
}