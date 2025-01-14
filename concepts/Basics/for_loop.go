package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 5; i++ {
		// if i%2 == 0 {
		// 	continue
		// }
		// fmt.Printf("%d ", i)
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	const typedHello string = "Hello, 世界" //typed
	var s string                          //typed
	s = typedHello
	fmt.Println(s)

	type MyString string // typed but defferent string type
	var m MyString
	// m = typedHello // Type error
	const hello = "Hello"
	m = hello // will convert to type MyString

	fmt.Println(m)
}
