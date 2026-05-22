package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ для подписи токена (в реальном проекте должен лежать в .env)
var SecretKey = []byte("super_secret_key_for_trade_system")

// GenerateToken создает JWT токен со встроенным ID пользователя и его ролью
func GenerateToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Токен живет 24 часа
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}
