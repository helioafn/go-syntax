// Listing 4.6: Variable scoping rules

package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // Avaliable throughout the package

func main() {
	year := 2018 // era and year are in scope

	switch month := rand.Intn(12) + 1; month { // era, year and month are in scope
	case 2:
		day := rand.Intn(28) + 1 // era, year, month and day are in scope
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 // It's a new day
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 // It's a new day as well
		fmt.Println(era, year, month, day)
	} // from here on month and day are out of scope
} // from here year is no longer in scope
