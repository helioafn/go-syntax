// Listing 4.5: Short declaration in a switch statement

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}

	// fmt.Println(num) // undefined: num

	// Quick check 4.2
	// num would be visible outside the scope of its declarations (if-statement and switch-case)
}
