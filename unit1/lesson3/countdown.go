// Listing 3.8: A countdown loop

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10

	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }

	// fmt.Println("Liftoff!")

	// Quick check 3.6
	for count > 0 {
		if rand.Intn(100) == 1 {
			break
		}
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}

	if count == 0 {
		fmt.Println("Blastoff!!")
	} else {
		fmt.Println("Launch failed!")
	}

}
