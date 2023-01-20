// calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	// declaration
	var a int
	var b int
	var c float32
	var d float32

	/* assigning variable
	===| method 1 |=== */
	a = 1
	b = 2
	c = 1.0
	d = 2.0 // = is just assigning value and the variable must be declared for the first time

	// ===| method 2 |===
	e := 3.0 // := is defining and assigning value, and its defined float64 by default type
	f := 4.0 // variable can't be declared becuase of :=

	// ===| method 3 |===
	g, h := 5.0, 6.0 // we can assign to two values at the same time

	// print command
	fmt.Printf("a=%v, type of %T\n", a, a) // %v is verb, v stands for value
	fmt.Printf("b=%v, type of %T\n", b, b) // %T is special verb, T stands for type
	fmt.Println("=========================")
	fmt.Printf("c=%v, type of %T\n", c, c)
	fmt.Printf("d=%v, type of %T\n", d, d)
	fmt.Println("=========================")
	fmt.Printf("e=%v, type of %T\n", e, e)
	fmt.Printf("f=%v, type of %T\n", f, f)
	fmt.Println("=========================")
	fmt.Printf("g=%v, type of %T\n", g, g)
	fmt.Printf("h=%v, type of %T\n", h, h)
	fmt.Println("=========================")

	// calculate means
	mean := (a + b) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
	mean1 := (c + d) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean1, mean1)
	mean2 := (e + f) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean2, mean2)
	mean3 := (g + h) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean3, mean3)
}
