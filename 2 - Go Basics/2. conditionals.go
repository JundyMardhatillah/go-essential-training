package main

import (
	"fmt"
)

func main() {
	// variable
	x := 10

	// if
	if x > 5 {
		fmt.Println("x is big")
	}

	// if else
	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	// logical operators to combine conditions
	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	// using or by specifying the two vertical bars
	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	// initialization if statement
	a := 11.0
	b := 12.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	// initialization switch statement
	n := 2

	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	// naked switch with conditions in the case statement
	m := 20

	switch {
	case m > 10:
		fmt.Println("m is very big")
	case m > 100:
		fmt.Println("m is big")
	default: // its called default clause
		fmt.Println("m is small")
	}
}
