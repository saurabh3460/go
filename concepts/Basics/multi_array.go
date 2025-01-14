package main

import "fmt"

// Multidimensional arrays
func printarray(a [3][2]string) {
	for _, s1 := range a {
		for _, s2 := range s1 {
			fmt.Printf("%s ", s2)
		}
		fmt.Printf("\n")
	}
}
func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printarray(a)

}
