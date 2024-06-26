package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "QuanNM-PTIT"

func GenerateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SecretKey))
}
