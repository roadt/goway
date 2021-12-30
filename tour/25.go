/*
Struct Fields

Struct fields are accessed using a dot.
*/

package main

import "fmt"

type Vec struct {
	x int
	y int
}

func main() {
	v := Vec{1,2}
	v.x = 4
	fmt.Println(v.x)
}