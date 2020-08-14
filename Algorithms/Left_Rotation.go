package main

import "fmt"

func rotLeft(a []int32, d int32) []int32 {
	temp := make([]int32, len(a))
	var value int32
	var i int32
	for i = 0; i < d; i++ {

		if i == 0 {
			value = a[0]
			temp = a[i+1:]
		} else {
			value = temp[0]
			temp = temp[1:]
		}
		temp = append(temp, value)
	}
	return temp
}

func main() {
	c := []int32{1, 2, 3, 4, 5}
	fmt.Println(rotLeft(c, 4))
}
