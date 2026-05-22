package management

import (
	"context"
	"errors"

	"trade-organization/internal/domain"
)

// StoreRepository — порт для CRUD операций над справочником точек
type StoreRepository interface {
	CreateStore(ctx context.Context, store domain.Store) (int, error)
	GetStoreByID(ctx context.Context, id int) (domain.Store, error)
}

type StoreService struct {
	repo StoreRepository
}

func NewStoreService(repo StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

// CreateStore содержит бизнес-правила создания новых точек
func (s *StoreService) CreateStore(ctx context.Context, store domain.Store) (int, error) {
	if store.Area <= 0 {
		return 0, errors.New("store area must be greater than zero")
	}
	if store.Type != "Department_Store" && store.Type != "Shop" && store.Type != "Kiosk" && store.Type != "Tray" {
		return 0, errors.New("invalid store type")
	}

	return s.repo.CreateStore(ctx, store)
}
