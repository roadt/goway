/*
Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & constructs a pointer to a struct literal.
*/

package main
import "fmt"

type vec struct {
	x,y int
}

var (
	p = vec{1,2} // has type vec
	q = &vec{1,2} // has type *vec
	r = vec{x: 1} // y:0 is implicit
	s = vec{} // x:0 and y:0
)

func main() {
	fmt.Println(p,q,r,s)
}
