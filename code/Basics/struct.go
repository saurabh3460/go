package main

import "fmt"

// named structure

type Employee struct {
	firstname, lastname string
	age                 int
}

// anonymous structure
var employee struct {
	firstname, lastname string
	age                 int
}

func main() {
	var emp1 Employee = Employee{
		firstname: "jack",
		lastname:  "rayen",
		age:       2,
	}
	emp1.firstname = ""
	fmt.Println("Hello", emp1.firstname)

}
