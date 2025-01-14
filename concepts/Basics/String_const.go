package main

import "fmt"

// String Constants

/*
Any value enclosed between double quotes is a string constant in Go.
For example strings like "Hello World" or "Sam" are all constants in Go.

A string constant like "Hello World" does not have any type.

*/
func main() {

	var hello = "hello" // it is untyped
	/*
		In the above case we have assigned "Hello World" to a named constant hello. Now does the constant hello have a type? The answer is No. The constant still doesn't have a type.
	*/

	fmt.Println(&hello)
}
