// 3.4 Logical Operators (|| -> OR, && -> AND)
package main

import "fmt"

func main() {
	var year = 2000
	fmt.Printf("The year is %v, should you leap?\n", year)

	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground")
	}
}
