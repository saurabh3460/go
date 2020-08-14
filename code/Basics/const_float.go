package main

import (
	"fmt"
)

func main() {
	type MyFloat64 float64
	const Zero = 0.0
	const TypedZero float64 = 0.0
	var mf MyFloat64
	mf = 0.0  // Ok
	mf = Zero // oK
	// mf = TypedZero //Bad
	fmt.Println(mf)

	var f32 float32
	f32 = 0.0  //untyped const float
	f32 = Zero // Zero is untyped const float
	// f32 = TypedZero // Bad: TypedZero is float64 not float32.
	fmt.Println(f32)
}
