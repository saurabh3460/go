package main

import "fmt"

type myInfo struct {
	name     string
	lastname string
}

/*--------------------6. Passing pointer to a function-------------------------*/
func greeting(message *string) {

	*message = "Hello " + *message + "!"
}

// func removeColor(indexI, indexJ int, ar []int32) {
// 	firstSlice1 := ar[:indexI]
// 	lastSlice1 := ar[indexI+1:]
// 	ar = append(firstSlice1, lastSlice1...)
// 	firstSlice2 := ar[:indexJ]
// 	lastSlice2 := ar[indexJ+1:]

// 	fmt.Printf("%v  %v\n", indexI, indexJ)
// 	ar = append(firstSlice2, lastSlice2...)
// 	fmt.Printf("%v \n", ar)

// }
func find(item1, item2 int, ar []int) bool {
	fmt.Printf(" for check %v \n", ar)
	for _, i := range ar {
		fmt.Printf("%v %v ", item1, item2)
		if item1 == i || item2 == i {

			fmt.Printf("will not %v %v ", item1 == i, item2 == i)
			fmt.Println(ar)
			return true
		}

	}
	fmt.Printf("%v %v ", item1, item2)

	// fmt.Println(ar)
	return false
}
func main() {

	// var m1 *myInfo = new(myInfo)
	// (*m1).lastname = "hdsdk"
	// m1.name = "idisjdi"
	// fmt.Println(*m1)

	// name := "Saurabh"
	// originalMessage := &name
	// greeting(originalMessage)
	// fmt.Println(*originalMessage)
	// fmt.Println(name)
	count := 0
	ar := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	var in []int
	for indexI, i := range ar {
		for indexJ, l := range ar {
			if i == l {
				if indexI == indexJ {
					continue
				}
				if !find(indexI, indexJ, in) {
					// fmt.Printf("%v %v \n", indexI, indexJ)
					in = append(in, indexI, indexJ)
					count++
					fmt.Println(in)
				}
			}

		}
	}
	fmt.Println(count)

}
