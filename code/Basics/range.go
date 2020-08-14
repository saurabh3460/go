package main

import "fmt"

func main() {
	a := []float32{212.2324, 4343.434, 4343}
	sum := float32(0)
	for _, v := range a { //range returns both the index and value
		sum += v
	}
	fmt.Println(sum)
}
