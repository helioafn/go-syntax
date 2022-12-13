// 9.4 Manipulating characters with Ceaser cipher

package main

import "fmt"

func main() {
	c := 'a'
	c = c + 3
	fmt.Printf("%c\n", c)

	// 9.4.1 - A modern variant
	// ROT13 (rotate 13) is a 20th century variant of Caesar cipher. It has one difference: it adds 13 instead of 3
	// Ciphering and deciphering are the same convenient operation
	message := "uv vagreangvbany fcnpr fgngvba"
	fmt.Println(len(message)) // len function return the length of a variety of types. In this case len returns the
	// length of a string in bytes

	// Experiment: Decipher the quote from /Julius Caesar: "L fdph, L vdz, L frqtxhuhg"
	quote := "L fdph, L vdz, L frqtxhuhg"
	for _, c := range quote {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c = c - 3
			if c > 'z' {
				c = c - 26
			}
		}

		fmt.Printf("%c", c)
	}

	fmt.Println()
}
