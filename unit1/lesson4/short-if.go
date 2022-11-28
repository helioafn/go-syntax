// Listing 4.4: Short declaration in a if statement

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// num 'lives' only in the scope of this if statement
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	// fmt.Println(num) // undefined: num
}
