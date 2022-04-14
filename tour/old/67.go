/*
Default Selection

The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:

select {
case i := <-c:
	// use i
default:
	// receiving from c would block
}
Note: This example won't run through the web-based tour user interface because the sandbox environment has no concept of time. You may want to install Go to see this example in action.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for{
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(5e7)
		}
	}
}
