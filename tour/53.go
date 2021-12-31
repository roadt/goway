/*
Interfaces

An interface type is defined by a set of methods.

A value of interface type can hold any value that implements those methods.
*/

package main

import (
	"fmt"
	"math"
)

type  Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f   // a my float implements Abser
	a = &v // a  *vertex implements Abser
//	a = v // a Vertex, does NOT implemnt Abser.   qr_2973JAU.go:27:4: error: incompatible types in assignment (method ‘Abs’ requires a pointer)
	
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64 (f)
}

type Vertex struct {
	x,y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y+v.y)
}

