package dto

import "time"

// SaleResponse — то, что видит фронтенд при просмотре чека
type SaleResponse struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"sale_date"`
	SellerID   int       `json:"seller_id"`
	CustomerID *int      `json:"customer_id,omitempty"`
}

// CreateSaleRequest — то, что приходит с фронтенда при оформлении продажи
type CreateSaleRequest struct {
	SellerID   int  `json:"seller_id" binding:"required"`
	CustomerID *int `json:"customer_id"`
	Items      []struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,gt=0"`
	} `json:"items" binding:"required,dive"`
}
