/*
>If we declare identifier in lowercase letter, it will be visible within the package only. But if we declare package in uppercase letter, it will be
 visible within and outside the package which is also known as exported.
> The dot . Operator is used to access the identifier e.g. pack.Age where pack is the package name and Age is the identifier.
> There are two variants of for loop in Go: Counter-controlled iteration and Condition-controlled iteration.
*/

package main // It defines a stand alone executable program, not a library
import (
	"fmt"
)

func main() {

	no := []int{43, 44, 54}
	for i, v := range no {
		fmt.Println(i)
		fmt.Println(v)
	}
}
