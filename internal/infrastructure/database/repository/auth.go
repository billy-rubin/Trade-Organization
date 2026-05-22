package repository

import (
	"context"
	"errors"
	"fmt"
	"trade-organization/internal/infrastructure/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"trade-organization/internal/domain"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `
		SELECT user_id, username, password_hash, role, seller_id 
		FROM App_Users 
		WHERE username = $1
	`

	var userDB database.UserDB

	err := r.pool.QueryRow(ctx, query, username).Scan(
		&userDB.ID,
		&userDB.Username,
		&userDB.PasswordHash,
		&userDB.Role,
		&userDB.SellerID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return userDB.ToDomain(), nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user domain.User) error {
	query := `
		INSERT INTO App_Users (username, password_hash, role, seller_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.pool.Exec(ctx, query, user.Username, user.PasswordHash, user.Role, user.SellerID)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}
