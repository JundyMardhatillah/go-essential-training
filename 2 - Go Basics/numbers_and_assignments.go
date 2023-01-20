// calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	// declaration
	var a int
	var b int

	// assigning variable
	a = 1
	b = 2

	// print command
	fmt.Printf("x=%v, type of %T\n", a, a) // %v is verb, v stands for value
	fmt.Printf("x=%v, type of %T\n", b, b) // %T is special verb, T stands for type

	// calculate means
	mean := (a + b) / 2.0 // == (1 + 2) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
