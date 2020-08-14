package main

import (
	"fmt"

	"github.com/saurabh3460/webserver/models"
)

func main() {
	// name := "Jack ryen"
	// fmt.Println(name)
	// codeName := &name
	// fmt.Println(*codeName) // show the value ate this 0xc0000481e0 address
	// type user struct {
	// 	ID        int
	// 	admin     bool
	// 	dinNo     float32
	// 	firstname string
	// }
	// var (
	// 	u user
	// )
	// fmt.Println(u)
	u := models.User{
		ID:        23,
		FirstName: "hell",
		LastName:  "go",
	}
	fmt.Println(u)
}
