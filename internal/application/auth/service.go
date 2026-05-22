package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"trade-organization/internal/domain"
	"trade-organization/pkg/jwt"
)

// AuthRepository — интерфейс, который должен быть реализован в слое инфраструктуры
type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

// AuthService управляет бизнес-логикой аутентификации
type AuthService struct {
	repo AuthRepository
}

// NewAuthService — конструктор сервиса
func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Login проверяет учетные данные и возвращает JWT-токен и роль пользователя
func (s *AuthService) Login(ctx context.Context, username, password string) (string, string, error) {
	// 1. Ищем пользователя в базе данных (через интерфейс)
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		// Возвращаем абстрактную ошибку, чтобы не подсказывать хакерам, что именно не так
		return "", "", errors.New("invalid username or password")
	}

	// 2. Сравниваем введенный пароль с хэшем из БД
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// 3. Генерируем JWT токен
	token, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, user.Role, nil
}
