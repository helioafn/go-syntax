// 7.1: Declaring integer variables

package main

func main() {
	// Five integer types are signed (can represent both positive and negative)

	// var year int = 2018
	var month uint = 2

	// When using type inference, Go will always pick the int type for a literal whole number
	year := 2018
	// var year = 2018
	// var year int = 2018

	// 7.1.1 - Integer types for every occasion
	// Integers have lots of sizes (how much it holds and memory needed)
	// There are eight architecture-independent types

	/*
		Type	Range	 								Storage
		int8	-128 to 127								8-bit (1 byte)
		uint8	0 to 255

		int16	-32,768 to 32,767						16-bit (2 bytes)
		uint16 	0 to 65535

		int32	-2,147,483,648 to						32-bit (4 bytes)
				2,147,483,647
		uint32 	0 to 4,294,967,295

		int64	-9,223,372,036,854,775,808 to
				9,223,372,036,854,775,807				64-bit (8 bytes)
		uint64 	0 to 18,446,744,073,709,551,615
	*/

	/*
		int and uint are optimal for the target device.
		For example: The Go playground, Raspberry Pi 2 and older mobile phones
		provide a 32-bit environment where both int/uint are 32-bit values.
		Any recent computer will provide a 64-bit environment where int/uint will be 64-bit values
	*/

}
