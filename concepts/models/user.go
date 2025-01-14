package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

/*
users []*User is slice of pointer User not fixed and efficient because we just need to kar
nextID = 1 we can assign in this way at package level

*/
