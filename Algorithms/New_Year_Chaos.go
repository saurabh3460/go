package main

import (
	"fmt"

	"sort"
)

func int32Toint(a []int32) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = int(v)
	}
	return b
}
func minimumBribes(q []int32) {
	b := int32Toint(q)
	sort.Ints(b)
	l := len(b) - 1
	count := 0
	tc := false
	for i := l; i >= 0; i-- {
		for j := l; j >= 0; j-- {
			if b[i] == int(q[j]) {
				fmt.Println(b[i], q[j], b[i] == int(q[j]), i, j, i-j)
				if (i - j) == 2 {
					count += 2
				} else if i-j == 1 {
					count++
				} else if (i - j) > 2 {
					fmt.Println("Too chaotic")
					tc = true
					i = -1
					break
				}
			}
		}
	}
	if !tc {
		fmt.Println(count)
	}

}

func main() {
	c := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(c)
}
