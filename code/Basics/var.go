package main

import "fmt"

func main() {

	var (
		name   = "saurabh"
		age    = "19"
		height int
	)
	var l int = 32
	var ll = 22 //// type will be inferred
	lo := "new"
	fmt.Println(name, age, height, l, ll, lo)

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

}
