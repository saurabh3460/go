package main

import "fmt"

func max(a, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}
func twoStacks(x int32, a []int32, b []int32) int32 {

	var sum1 int32 = a[0]
	var sum2 int32 = b[0]
	var count1 int32 = 0
	var count2 int32 = 0
	a = a[1:]
	b = b[1:]
	var sum int32 = 0
	var count int32 = 0

	al := int32(len(a))
	bl := int32(len(b))

	l := max(al, bl)
	fmt.Println("lenght", l, x)
	var i int32

	for i = 0; i < l; i++ {
		if i <= al-1 {
			if sum1+a[0] <= x {
				sum1 += a[0]
				a = a[1:]
				fmt.Println("in a", sum1)
				count1++
			} else {
				a = a[1:]
			}
		}
		if i <= bl-1 {
			if sum2+b[0] <= x {
				sum2 += b[0]
				b = b[1:]
				// fmt.Println("in b", sum, i, b[0])
				count2++
			} else {
				b = b[1:]
				fmt.Println("in b in else", b)
			}
		}
	}
	return count1 + count2
}

func main() {

	const x int32 = 77
	// ar := []int32{17, 5, 0}
	// br := []int32{10, 8, 2, 1, 13, 1, 14, 18, 9, 18, 16, 19, 5}
	// ar := []int32{4, 2, 4, 6, 1}
	// br := []int32{2, 1, 8, 5}

	ar := []int32{11, 4, 1, 11, 10, 6, 5, 1, 1, 19, 9, 7, 1, 16, 14, 2, 14, 19, 0, 4, 10, 1, 1, 11, 2, 11, 5, 7, 0, 13}
	br := []int32{16, 4, 17, 9, 15, 19, 16, 0, 0, 9, 11, 10, 17, 4, 18, 3, 6}

	a := ar[:]
	b := br[:]
	fmt.Println("count", twoStacks(x, a, b))
}
