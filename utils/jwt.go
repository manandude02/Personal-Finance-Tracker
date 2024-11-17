package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT generates a new JWT for a given user ID.
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(jwtSecret)
}
