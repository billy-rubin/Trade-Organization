package trading

import (
	"context"
	"errors"
	"fmt"

	"trade-organization/internal/domain"
)

type TradeRepository interface {
	CreateSale(ctx context.Context, sale domain.Sale, details []domain.SaleDetail) (int, error)
	TransferProduct(ctx context.Context, fromStoreID, toStoreID, productID, quantity int) error
}

type TradeService struct {
	repo TradeRepository
}

func NewTradeService(repo TradeRepository) *TradeService {
	return &TradeService{repo: repo}
}

func (s *TradeService) CreateSale(ctx context.Context, sale domain.Sale, details []domain.SaleDetail) (int, error) {
	if len(details) == 0 {
		return 0, errors.New("sale must contain at least one item")
	}

	for _, item := range details {
		if item.Quantity <= 0 {
			return 0, fmt.Errorf("invalid quantity for product ID %d: must be greater than zero", item.ProductID)
		}
	}

	return s.repo.CreateSale(ctx, sale, details)
}

func (s *TradeService) TransferProduct(ctx context.Context, fromStoreID, toStoreID, productID, quantity int) error {
	if quantity <= 0 {
		return errors.New("transfer quantity must be greater than zero")
	}
	if fromStoreID == toStoreID {
		return errors.New("source and destination stores cannot be the same")
	}

	return s.repo.TransferProduct(ctx, fromStoreID, toStoreID, productID, quantity)
}
