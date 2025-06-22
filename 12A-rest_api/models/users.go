package models

import (
	"errors"
	"example.com/rest_api/db"
	"example.com/rest_api/utils"
)

func (u User) Save() error {
	query := "INSERT INTO users(email, password) values (? ,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return errors.New("Credentials invalid!")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievePassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid!")
	}

	return nil
}
