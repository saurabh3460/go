package main

import "fmt"

func main() {
	name := "saurabh"
	for _, i := range name {
		defer fmt.Printf("%c", i)
	}
}
