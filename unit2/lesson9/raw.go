// 9.1 - Declaring string variables
package main

import "fmt"

func main() {
	peace := "peace"
	// var peace = "peace"
	// var peace string = "peace"

	var blank string // zero value for string is "" (empty string)

	fmt.Println(peace)
	fmt.Println(blank)

	// Listing 9.1 - Raw string literals

	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can span multiple lines with the \n escape sequence`)

	// Raw string literals can span multiple lines of source code

	// Listing 9.2 - Muliple-line raw string literals
	fmt.Println(`
		peace be upon you
		upon you be peace`)

	// Listing 9.3 - String type
	fmt.Printf("%v is %[1]T\n", "literal string")       // literal string is a string
	fmt.Printf("%v is a %[1]T\n", `raw string literal`) // raw string literal is a string

	// Quick check 9.1
	// A: A raw string literal `C:\go`

}
