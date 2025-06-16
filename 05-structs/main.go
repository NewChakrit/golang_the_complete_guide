package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//user := User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	BirthDate: birthdate,
	//	CreatedAt: time.Now(),
	//}

	user := User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}

	user.outputUserDetails()

	//outputUserDetails(&user)
}

func (u User) outputUserDetails() {

	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
