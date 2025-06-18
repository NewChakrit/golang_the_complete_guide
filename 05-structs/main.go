package main

import (
	"fmt"
	"structs_and_custom_types/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	email := getUserData("Please enter your email: ")
	password := getUserData("Please enter your password: ")

	//appUser := User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	BirthDate: birthdate,
	//	CreatedAt: time.Now(),
	// }

	//var appUser *user.User

	var appUser *user.User

	//appUser, err := user.NewUser(firstName, lastName, birthdate)
	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}

	admin := user.NewAdmin(email, password)
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	//outputUserDetails(&appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	//fmt.Scan(&value)
	fmt.Scanln(&value) // If pass enter, It will finish push value.
	return value
}
