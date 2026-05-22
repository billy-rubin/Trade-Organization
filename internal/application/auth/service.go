package auth

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"trade-organization/internal/domain"
	"trade-organization/pkg/jwt"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		// Возвращаем абстрактную ошибку, чтобы не подсказывать хакерам, что именно не так
		return "", "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid username or password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, user.Role, nil
}

func (s *AuthService) Register(ctx context.Context, username, password, role string, sellerID *int) error {
	_, err := s.repo.GetUserByUsername(ctx, username)
	if err == nil {
		return errors.New("user with this username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		SellerID:     sellerID,
	}

	return s.repo.CreateUser(ctx, user)
}
