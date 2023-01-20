package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	// how many bytes are in the book
	fmt.Println(len(book)) // use the built-in len function

	// I can access the first bite with the square brackets and index zero
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// to get a part of a string, I can use slice [start:end]
	fmt.Println(book[4:11]) // from 4 up to 11 not including 11. it will be 'colour'

	/* in slice I can also skip the end or skip the beginning
	slice (no end) */
	fmt.Println(book[4:]) // it will be 'colour of magic'

	// slice (no start)
	fmt.Println(book[:4]) // it will be 'The'

	// use + to concatenate strings
	fmt.Println("t" + book[1:]) // it will change from 'T' to 't'

	// unicode
	fmt.Println("It was ½ price!") // it will be as I expected

	// multi line
	poem := `
		The road goes ever on
		Down from the door where it began
		...
		` // finally I can use Raw strings starts with a backtick
	fmt.Println(poem)
}
