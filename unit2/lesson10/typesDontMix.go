// 10.1 Types don't mix

package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := "Launch in T minus " + "10 seconds" // plus (+) operator join string together
	timer := strconv.Itoa(10)                        // Converting a int to a string
	amazingNumber, _ := strconv.Atoi("1337")         // Converting a string to int
	fmt.Println(amazingNumber)
	// countdown = "Launch in T minus " + 10 + " seconds." Invalid operation: mismatched types string and int

	countdown = "Launch in T minus " + timer + " seconds."

	fmt.Println(countdown)

	age := 25
	marsDays := 687
	earthDays := 365.2425

	//  invalid operation: mismatched types int and float64
	// fmt.Println("I am", age*earthDays/marsDays, "years old on Mars.")

	// Quick check 10.1 - A: is an invalid operation (mismatched types: string and int)
}
