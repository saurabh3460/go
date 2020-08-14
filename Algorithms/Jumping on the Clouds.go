package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	var ci int32
	var jump int32
	l := len(c) - 1
	for index, val := range c {
		if index < int(ci) {
			continue
		}
		if val == 0 {
			if index+1 <= l && c[index+1] == 0 {
				if index+2 <= l && c[index+2] == 0 {
					ci = int32(index + 2)
					jump++
				} else {
					ci = int32(index + 1)
					jump++
				}
			}
		} else {
			if index+1 <= l && c[index+1] == 0 {
				ci = int32(index + 1)
				jump++
			}
		}
	}
	return jump
}

func main() {
	cloud := []int32{0, 0, 0, 1, 0, 0}
	result := jumpingOnClouds(cloud)
	fmt.Println(result)
}
