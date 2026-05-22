package supply

import (
	"context"
	"errors"

	"trade-organization/internal/domain"
)

type SupplyRepository interface {
	CreateRequest(ctx context.Context, req domain.Request, details []domain.RequestDetail) (int, error)
	GenerateOrderFromRequest(ctx context.Context, requestID int, supplierID int) error
	ReceiveOrder(ctx context.Context, orderID int, storeID int) error
	GetOrderDetails(ctx context.Context, orderID int) ([]domain.OrderDetailReport, error)
}

type SupplyService struct {
	repo SupplyRepository
}

func NewSupplyService(repo SupplyRepository) *SupplyService {
	return &SupplyService{repo: repo}
}

func (s *SupplyService) CreateRequest(ctx context.Context, req domain.Request, details []domain.RequestDetail) (int, error) {
	if len(details) == 0 {
		return 0, errors.New("request must contain at least one item")
	}

	for _, item := range details {
		if item.Quantity <= 0 {
			return 0, errors.New("quantity of requested items must be positive")
		}
	}

	req.Status = "New"
	return s.repo.CreateRequest(ctx, req, details)
}

func (s *SupplyService) ProcessRequestToOrder(ctx context.Context, requestID int, supplierID int) error {
	if requestID <= 0 || supplierID <= 0 {
		return errors.New("invalid request or supplier ID")
	}
	return s.repo.GenerateOrderFromRequest(ctx, requestID, supplierID)
}

func (s *SupplyService) ReceiveOrder(ctx context.Context, orderID int, storeID int) error {
	if orderID <= 0 || storeID <= 0 {
		return errors.New("invalid order or store ID")
	}
	return s.repo.ReceiveOrder(ctx, orderID, storeID)
}

func (s *SupplyService) GetOrderDetails(ctx context.Context, orderID int) ([]domain.OrderDetailReport, error) {
	return s.repo.GetOrderDetails(ctx, orderID)
}
