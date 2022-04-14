/*
The new function

The expression new(T) allocates a zeroed T value and returns a pointer to it.

var t *T = new(T)
or

t := new(T)
*/
package main

import "fmt"

type vec struct {
	x,y int
}

func main() {
	v:= new(vec)
	fmt.Println(v)
	v.x, v.y = 11, 9
	fmt.Println(v)
}