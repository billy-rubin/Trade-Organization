package dto

import (
	"time"

	"trade-organization/internal/domain"
)

type CreateRequestItem struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,gt=0"`
}

type CreateSupplyRequestPayload struct {
	StoreID int                 `json:"store_id" binding:"required"`
	Items   []CreateRequestItem `json:"items" binding:"required,dive"`
}

func (req *CreateSupplyRequestPayload) ToDomain() (domain.Request, []domain.RequestDetail) {
	request := domain.Request{
		Date:    time.Now(),
		StoreID: req.StoreID,
	}

	details := make([]domain.RequestDetail, 0, len(req.Items))
	for _, item := range req.Items {
		details = append(details, domain.RequestDetail{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return request, details
}

type GenerateOrderPayload struct {
	RequestID  int `json:"request_id" binding:"required"`
	SupplierID int `json:"supplier_id" binding:"required"`
}

type ReceiveOrderPayload struct {
	OrderID int `json:"order_id" binding:"required"`
	StoreID int `json:"store_id" binding:"required"`
}
