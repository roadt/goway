/*
Pointers

Go has pointers, but no pointer arithmetic.

Struct fields can be accessed through a struct pointer. The indirection through the pointer is transparent.
*/
package main

import "fmt"

type  vec struct {
	x int
	y int
}

func main() {
	p := vec{1,2}
	q := &p
	q.x = 1e9
	fmt.Println(p)
}