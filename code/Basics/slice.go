package main

import (
	"fmt"
)

// Slices do not own any data on their own
func main() {

	a := [5]int{32, 32, 324, 3, 41}
	var b []int = a[1:5]
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(b)
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
}
