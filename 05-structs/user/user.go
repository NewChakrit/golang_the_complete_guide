package user

import (
	"errors"
	"fmt"
	"time"
)

type Admin struct {
	Email    string
	Password string
	User
}

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
	//u.BirthDate = ""
}

func NewAdmin(email, password string) Admin {

	return Admin{
		Email:    email,
		Password: password,
		User: User{
			"ADMIN",
			"ADMIN",
			"----",
			time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate, u.BirthDate)
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName or lastName or birthdate are required")
	}
	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}
