package main

import (
	"fmt"
	"time"
)

type user struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func (u *user) clearUserName() {
	u.FirstName = ""
	u.LastName = ""
	u.BirthDate = ""
}

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//appUser := User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	BirthDate: birthdate,
	//	CreatedAt: time.Now(),
	// }

	var appUser *user
	appUser = newUser(firstName, lastName, birthdate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

	//outputUserDetails(&appUser)
}

func (u user) outputUserDetails() {

	fmt.Println(u.FirstName, u.LastName, u.BirthDate, u.BirthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
