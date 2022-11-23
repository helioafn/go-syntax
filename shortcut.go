// 2.4.2 Increment and assignment operators

package main

func main() {
	// Listing 2.5
	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783 // is equivalent as above

	// Increment operator
	var age = 25
	age = age + 1 // Happy birthday :D
	age += 1      // is equivalent as the code above
	age++         // is equivalent as the code above

	// You can decrement using count-- or other assignment operators like: /=, *=

	// Quick check 2.5
	weight -= 2
}
