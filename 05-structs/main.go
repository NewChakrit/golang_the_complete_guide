package main

import (
	"errors"
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
	//u.BirthDate = ""
}

func (u *user) outputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate, u.BirthDate)
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName or lastName or birthdate are required")
	}
	return &user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
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
	appUser, err := newUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

	//outputUserDetails(&appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	//fmt.Scan(&value)
	fmt.Scanln(&value) // If pass enter, It will finish push value.
	return value
}
