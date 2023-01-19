// ====| anatomy of a Go program |====
package main // package main is a special package that will make your code compile to an executable and its not library. All codes in Golang should be in the package.

import (
	"fmt" // fmt stands for the format package.
)

/*
func main() is a special type of function, it's the entry point of the executable programs,
and it's the first function that is executed when a Go program is run.
The main function is defined with the keyword func and its name is always main.
*/
func main() {
	fmt.Println("Welcome Gophers")
	/* fmt is a package that provides functions for formatting and printing text.
	The package includes functions like Println, Printf, Sprint, Sprintf, among others.
	The functions are used to output text to the console, format text, and manipulate strings */
}
