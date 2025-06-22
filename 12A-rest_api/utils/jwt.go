package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "supersecret" // key ที่ ใช้ sign token อารมณ์ลายเซ็น

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secretKey)) // todo ต้องส่งเป็น []byte()
	// ผลลัพธ์เป็นคล้ายๆ random string ยาวๆ
}

/* todo Note :
เวลาปัจจุบัน: 2025-06-22 20:05:23.123456789 +0700 +0700
Unix timestamp: 1719061523
ประเภทข้อมูลของ Unix timestamp: int64
*/

func VerifyToken(token string) error {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // jwt.SigningMethodHMAC คือ type เช็ค
		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse token.")
	}

	tokenIsValid := parseToken.Valid
	if !tokenIsValid {
		return errors.New("Invalid token!")
	}

	//claims, ok := parseToken.Claims.(jwt.MapClaims)
	//if !ok {
	//	return errors.New("Invalid token claims.")
	//}

	//email := claims["email"].(string)  // .(string) คือ func เอาไว้เช็คว่าเป็น string มี return ok แต่เราไม่สน
	//userID := claims["userID"].(int64) //  .(int64) คือ func เอาไว้เช็คว่าเป็น int64

	return nil
}
