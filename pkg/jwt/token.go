package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("super_secret_key_for_trade_system")

func GenerateToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}
