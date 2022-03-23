/*

Exercise: Maps

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _,v := range strings.Fields(s) {
		ret[v]  = ret[v] + 1
	}
	return ret
}


func main() {
	wc.Test(WordCount)
}

