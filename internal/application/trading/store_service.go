package trading

import (
	"context"
	"trade-organization/internal/domain"
)

type StoreRepository interface {
	Create(ctx context.Context, store domain.Store) (domain.Store, error)
	GetByID(ctx context.Context, id int) (domain.Store, error)
}

type StoreService struct {
	repo StoreRepository
}

func NewStoreService(repo StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) CreateStore(ctx context.Context, store domain.Store) (domain.Store, error) {
	return s.repo.Create(ctx, store)
}
