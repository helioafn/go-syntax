// 9.3 Pulling the Strings
package main

import "fmt"

func main() {
	// Listing 9.3
	message := "shalom"
	//c := message[5]
	// fmt.Printf("%c\n", c)

	// Quick check 9.3: Write a program to print each byte (ASCII character) of "shalom", one character per line
	for i := 0; i < 6; i++ {
		fmt.Printf("%c\n", message[i])
	}
}
