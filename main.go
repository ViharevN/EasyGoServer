package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	u := User{
		FirstName: "Misha",
		LastName:  "Popov",
	}
	fmt.Println(u.FullName())
}
