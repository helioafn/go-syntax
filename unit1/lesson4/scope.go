// Listing 4.1: Scoping rules

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	// In Go scopes tend to begin and end along the lines of curly braces {}.
	for count < 10 {
		var num = rand.Intn(10) + 1 // Visible only inside this for-loop
		fmt.Println(num)
		count++
	}
	// Quick check 4.1
	// 		1. You can re-use an variables names and can focus on the current scope
	// 		2. When you try to access an variable out of scope, it generates an error.
	// fmt.Println(num) // Generate an error (undefined: num)
	fmt.Println(count) // It works because count is in the main function scope
}
