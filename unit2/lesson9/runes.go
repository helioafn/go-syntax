// 9.2 - Characters, code points, runes and bytes

package main

import "fmt"

func main() {
	// To represent an single Unicode code point, Go provides rune (alias for int32)
	// byte is an alias for uint8, it's intended for binary data, although it can be used for chars of ASCII table

	// Type aliases
	// from Go 1.9+ You can declare your own type aliases (byte and rune came from the beginning of Go)
	type byte = uint8
	type rune = int32

	// Listing 9.4 - rune and byte
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) // 960 940 969 33

	// To display the characters rather than their numeric values, the %c format verb can be used with Printf()
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang) // πάω!

	// For a character literal (ex.: 'H'), if no type is specified, Go will infer rune
	grade := 'A'
	// var grade = 'A'
	// var grade rune = 'A' // This variable still contains a numeric value, in this case 65, the code point for 'A'.
	fmt.Println(grade)

	// Character literal can also be used with byte alias
	star := '*'
	fmt.Printf("%c\n", star)
	fmt.Println('', 'é')
	// Quick check 9.2
	// A1: 128 characters. A2: byte is an alias for uint8, and rune is an alias for int32.
	// A3: 42, 61514, 233
}
