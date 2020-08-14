package main

import "fmt"

func main() {
	var v int32
	var lvl int
	s := "UDDDUDUU"
	for _, i := range s {
		if string(i) == "U" {
			lvl++
		}
		if string(i) == "D" {
			lvl--
		}
		if lvl == 0 && string(i) == "U" {
			v++
		}

	}
	fmt.Println(v)
}
