package main

import "fmt"

func main() {
	// Printing on console is done through fmt package.
	fmt.Println("hello there!")

	// String concatenation works too
	fmt.Println("strings" + " are " + "concatenated")

	// Regular numeric math operations
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Basic boolean operations
	fmt.Println("true && false", true && false)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)
	fmt.Println("!!true", !!true)
}
