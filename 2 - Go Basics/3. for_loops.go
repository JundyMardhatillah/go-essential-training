package main

import (
	"fmt"
)

func main() {
	fmt.Println("====| traditional loop |====")
	for i := 0; i < 3; i++ {
		fmt.Println(i) // it will be 0 until 2
	}

	fmt.Println("====| traditional loop |====")
	for i := 0; i <= 3; i++ {
		fmt.Println(i) // it will be 0 until 3
	}

	fmt.Println("====| break loop in the middle |====")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break // stop until 1
		}
		fmt.Println(i) // it will be 0 until 1
	}

	fmt.Println("====| continue loop |====")
	// continue loop in the middle
	for i := 0; i <= 3; i++ {
		if i < 1 {
			continue // means go to the next loop until 1 instead of 0
		}
		fmt.Println(i) // it will be 3 until 1
	}

	fmt.Println("====| manually increment inside the for loop |====")
	a := 0
	for a < 3 {
		fmt.Println(a) // it will be 0 until 2
		a++ // increment +1
	}

	fmt.Println("====| while loop or infinite loop |====")
	b := 0
	for {
		if b > 2 {
			break // use manual break to exit out of it
		}
		fmt.Println(b) // it will be 0 until 2
		b++
	}
}
