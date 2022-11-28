// 2.5 Think of a number
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Listing 2.6
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	// Quick check 2.6
	var randomDistance = rand.Intn(401000000-56000000) + 56000001
	fmt.Println(randomDistance)

	// Extra
	const hoursPerDay = 24
	var days = 28
	var speedNeeded = 56000000 / (days * hoursPerDay) // km/h
	fmt.Println(speedNeeded)
}
