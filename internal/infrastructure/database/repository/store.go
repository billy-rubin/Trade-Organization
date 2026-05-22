package repository

import (
	"context"
	"fmt"

	"trade-organization/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreRepository struct {
	pool *pgxpool.Pool
}

func NewStoreRepository(pool *pgxpool.Pool) *StoreRepository {
	return &StoreRepository{pool: pool}
}

func (r *StoreRepository) CreateStore(ctx context.Context, store domain.Store) (int, error) {
	query := `
		INSERT INTO Store (area, counter_count, store_type)
		VALUES ($1, $2, $3)
		RETURNING store_id
	`
	var storeID int
	err := r.pool.QueryRow(ctx, query, store.Area, store.CounterCount, store.Type).Scan(&storeID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert store: %w", err)
	}
	return storeID, nil
}

func (r *StoreRepository) GetStoreByID(ctx context.Context, id int) (domain.Store, error) {
	return domain.Store{}, nil
}
