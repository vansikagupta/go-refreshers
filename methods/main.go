package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
}

// ShowUser is a function that does not bind to a named type
func ShowUser(user User) {
	fmt.Println(user.FirstName, user.LastName)
}

// GreetUser is a method associated to User type; it has value receiver
func (user User) GreetUser() {
	fmt.Println(user.FirstName, user.LastName)
}

// UpdateUser is also a method associated to User type; it has pointer receiver
func (user *User) UpdateUser(newLastName string) {
	user.LastName = newLastName // modifies the value of User
}

func main() {
	u := User{"Ron", "Miller"}
	ShowUser(u)               // calling a function
	u.GreetUser()             // calling method with value receiver
	(&u).UpdateUser("Mullen") // calling method with pointer receiver
	// let's see if the update reflects
	fmt.Println(u.FirstName, u.LastName)
}
