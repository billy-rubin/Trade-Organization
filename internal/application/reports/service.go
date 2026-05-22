package reports

import (
	"context"
	"errors"
	"time"

	"trade-organization/internal/domain"
)

type ReportRepository interface {
	GetStoreInventory(ctx context.Context, storeType string) ([]domain.InventoryReport, error)
	GetSellerOutput(ctx context.Context, sellerID int, startDate, endDate time.Time) (domain.SellerOutputReport, error)
	GetStoreProfitability(ctx context.Context, storeID int, startDate, endDate time.Time) (domain.ProfitabilityReport, error)
	GetTopActiveCustomers(ctx context.Context, limit int) ([]domain.ActiveCustomerReport, error)
	GetTurnoverByStoreType(ctx context.Context) ([]domain.TurnoverReport, error)
	GetStoreEfficiency(ctx context.Context, storeID int) (domain.StoreEfficiencyReport, error)
	GetSupplierDeliveries(ctx context.Context, supplierID, productID int, startDate, endDate time.Time) ([]domain.SupplierDeliveryReport, error)
	GetProductCustomers(ctx context.Context, productID, storeID int) ([]domain.ProductCustomerReport, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetInventory(ctx context.Context, storeType string) ([]domain.InventoryReport, error) {
	if storeType == "" {
		return nil, errors.New("store type cannot be empty")
	}
	return s.repo.GetStoreInventory(ctx, storeType)
}

func (s *ReportService) GetSellerOutput(ctx context.Context, sellerID int, startDate, endDate time.Time) (domain.SellerOutputReport, error) {
	if startDate.After(endDate) {
		return domain.SellerOutputReport{}, errors.New("start date cannot be after end date")
	}
	return s.repo.GetSellerOutput(ctx, sellerID, startDate, endDate)
}

func (s *ReportService) GetProfitability(ctx context.Context, storeID int, startDate, endDate time.Time) (domain.ProfitabilityReport, error) {
	if startDate.After(endDate) {
		return domain.ProfitabilityReport{}, errors.New("start date cannot be after end date")
	}
	return s.repo.GetStoreProfitability(ctx, storeID, startDate, endDate)
}

func (s *ReportService) GetTopCustomers(ctx context.Context, limit int) ([]domain.ActiveCustomerReport, error) {
	if limit <= 0 {
		limit = 5
	}
	if limit > 100 {
		limit = 100
	}
	return s.repo.GetTopActiveCustomers(ctx, limit)
}

func (s *ReportService) GetTurnover(ctx context.Context) ([]domain.TurnoverReport, error) {
	return s.repo.GetTurnoverByStoreType(ctx)
}

func (s *ReportService) GetStoreEfficiency(ctx context.Context, storeID int) (domain.StoreEfficiencyReport, error) {
	return s.repo.GetStoreEfficiency(ctx, storeID)
}

func (s *ReportService) GetSupplierDeliveries(ctx context.Context, supplierID, productID int, startDate, endDate time.Time) ([]domain.SupplierDeliveryReport, error) {
	return s.repo.GetSupplierDeliveries(ctx, supplierID, productID, startDate, endDate)
}

func (s *ReportService) GetProductCustomers(ctx context.Context, productID, storeID int) ([]domain.ProductCustomerReport, error) {
	return s.repo.GetProductCustomers(ctx, productID, storeID)
}
