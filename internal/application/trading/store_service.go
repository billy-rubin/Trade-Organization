package trading

import (
	"context"
	"trade-organization/internal/domain"
)

// StoreRepository интерфейс (реализован в infrastructure/database)
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

// CreateStore работает только с чистыми сущностями
func (s *StoreService) CreateStore(ctx context.Context, store domain.Store) (domain.Store, error) {
	// Здесь может быть бизнес-логика: проверки, расчеты и т.д.
	return s.repo.Create(ctx, store)
}
