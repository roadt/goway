/*
Errors

An error is anything that can describe itself as an error string. 
The idea is captured by the predefined, built-in interface type, error, with its single method, Error, returning a string:

type error interface {
	Error() string
}
	
The fmt package's various print routines automatically know to call the method when asked to print an error.
*/

package main

import (
	"fmt"  
	"time" 
)

type myerror struct {
	when time.Time
	what  string 
}

func (e *myerror) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run () error {
	return &myerror{
		time.Now(),
		"it didnt' work",
	}
}


func main() {
	if err := run(); err!=nil {
		fmt.Println(err)
	}
}
