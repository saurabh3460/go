package main

import (
	"fmt"
)

func main() {

	if num := 10; num%2 == 0 {
		fmt.Println("this is even")
	} else {
		fmt.Println("odd")
	}
}
