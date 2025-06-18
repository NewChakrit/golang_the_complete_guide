package user

import (
	"errors"
	"fmt"
	"time"
)

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
