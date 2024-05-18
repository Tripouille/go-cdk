package types

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user User) (string, error) {
	now := time.Now()
	validUntil := now.Add(time.Hour * 1).Unix(
)
	claims := jwt.MapClaims{
		"username": user.Username,
		"expires":  validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
