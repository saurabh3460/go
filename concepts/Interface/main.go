package main

import (
	"fmt"
)

type getName interface {
	getFirstName() string
	getLastName() string
}

type EmpInfo struct {
	firstName string
	lastName  string
	fullName  string
}

func (emp EmpInfo) getFirstName() string {
	return emp.firstName
}

func (emp EmpInfo) getLastName() string {
	return emp.lastName
}

func (emp EmpInfo) getFullName() string {
	return emp.fullName
}

func main() {
	fmt.Println("hello")
	var emp1 = EmpInfo{"saurabh", "shanu", "saurabh shanu"}
	var emp2 getName
	emp2 = emp1
	fmt.Println(emp2)

	// fmt.Println(emp1.getFirstName())
	// fmt.Println(emp1.getLastName())

}
