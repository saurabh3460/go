package main

import (
	"fmt"
)

func main() {

	const typedHello string = "Hello, 世界" //typed
	var s string                          //typed
	s = typedHello
	fmt.Println(s)

	type MyString string // typed but defferent string type mYsT
	var m MyString       // has MyString type
	// m = typedHello // Type error so typedHello and MyString both are deffrent
	const hello = "Hello" // untyped string constant
	m = hello             // will convert to type MyString

	fmt.Println(m)
	fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")
	fmt.Printf("%T: %v\n", hello, hello)
	/*
		because, unlike the typed constants {typedHello}, the untyped constants "Hello, 世界" and hello have no type.
		Assigning them to a variable of any type compatible with strings works without error.

		These untyped string constants are strings, of course, so they can only be used where a string is allowed,
		but they do not have type string.
	*/

	// Confusion Part
	str := "Hello, 世界"
	str = "2232e2"
	fmt.Println(str)
	
	/*
		and by now you might be asking, "if the constant is untyped, how does str get a type in this variable
		 declaration?" The answer is that an untyped constant has a default type, an implicit type that it
		  transfers to a value if a type is needed where none is provided. For untyped string constants,
		   that default type is obviously string,

		   In summary, a typed constant obeys all the rules of typed values in Go. On the other hand, an untyped constant does not carry a Go type in the same way and can be mixed and matched more freely. It does, however, have a default type that is exposed when, and only when, no other type information is available.

	*/
	// Booleans
	type MyBool bool
	const True = true
	const TypedTrue bool = true
	var mb MyBool
	mb = true // OK
	mb = True // OK
	//mb = TypedTrue // Bad
	fmt.Println(mb)

}
