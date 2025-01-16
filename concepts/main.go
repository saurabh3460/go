package main

import (
	structs "concepts/Structs"
	"fmt"
)

func main() {
	ep := structs.NewSlackProfile("John Doe", "johndoe", "Software Engineer", "1234567890", "johndoe.jpg", "/images/johndoe.jpg")
	fmt.Println("New Slack Profile:", ep)
	ep1 := structs.NewSlackProfile("John Doe", "johndoe", "Software Engineer", "1234567890", "johndoe.jpg", "/images/johndoe.jpg")

	fmt.Println("CheckDuplicateProfile:", ep.CheckDuplicateProfile(*ep1))
}
