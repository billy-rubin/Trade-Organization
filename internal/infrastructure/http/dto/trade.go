package dto

import (
	"time"
	"trade-organization/internal/domain"
)

type CreateSaleItemRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,gt=0"`
}

type CreateSaleRequest struct {
	CustomerID *int                    `json:"customer_id"`
	Items      []CreateSaleItemRequest `json:"items" binding:"required,dive"`
}

func (req *CreateSaleRequest) ToDomain(sellerID int) (domain.Sale, []domain.SaleDetail) {
	sale := domain.Sale{
		Date:       time.Now(),
		SellerID:   sellerID,
		CustomerID: req.CustomerID,
	}

	details := make([]domain.SaleDetail, 0, len(req.Items))
	for _, item := range req.Items {
		details = append(details, domain.SaleDetail{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return sale, details
}

type TransferProductRequest struct {
	FromStoreID int `json:"from_store_id" binding:"required"`
	ToStoreID   int `json:"to_store_id" binding:"required"`
	ProductID   int `json:"product_id" binding:"required"`
	Quantity    int `json:"quantity" binding:"required,gt=0"`
}
