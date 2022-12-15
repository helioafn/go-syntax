// Capstone: The Vigenère Cipher
package main

import "fmt"

func main() {
	fmt.Println("Hey!")
	cipheredMessage := "LFDPHLVDZLFRQTXHUHG"

	// Common Caesar cipher
	for _, c := range cipheredMessage {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c = c - 3
			if c > 'z' || c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
