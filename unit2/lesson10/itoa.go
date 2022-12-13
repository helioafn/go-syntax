// Listing 10.4 - Integer to ASCII
package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10
	// Itoa means Integer to ASCII. Unicode is a superset of the ASCII standard. The first 128 code points are the same,
	// that includes digits, English letters and common punctuation.
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	// presenting fmt.Sprintf(). It returns a string rather than displaying it.
	countdown = 9
	str = fmt.Sprintf("Launch in T minus %v seconds", countdown)
	fmt.Println(str)

	// strconv package provides Atoi function (ASCII to integer). A string may contain gibberish or a number that is too big.
	// So Atoi function may return a error
	cd, err := strconv.Atoi("10")
	if err != nil { // A nil value for err, indicates that no error occurred and everything is A-OK.
		// Oops, something went wrong
	} else {
		fmt.Println(cd)
	}

	// Quick check 10.4
	// A: Itoa from strconv package and Sprintf from fmt package

}
