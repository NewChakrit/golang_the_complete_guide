package utils

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	//if err != nil {
	//	return false
	//}
	//
	//return true

	return err == nil // todo สุดยอดดด!!
}
