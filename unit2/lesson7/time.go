// Listing 7.5 - 64-bit integers example (using time package)

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	future := time.Unix(math.MaxInt64-1, 0)
	fmt.Println(future)
	future = time.Unix(12622780800, 0)
	fmt.Println(future)

	// Quick check 7.7
	// A: a large enough to hold the desired values
}
