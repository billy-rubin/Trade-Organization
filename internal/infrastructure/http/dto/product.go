package dto

import "trade-organization/internal/domain"

type ProductResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func ToProductResponse(p domain.Product) ProductResponse {
	return ProductResponse{ID: p.ID, Name: p.Name, CurrentPrice: p.CurrentPrice}
}

type CreateProductDTO struct {
	Name         string  `json:"name" binding:"required"`
	CurrentPrice float64 `json:"current_price" binding:"required,gte=0"`
}

func (dto *CreateProductDTO) ToDomain() domain.Product {
	return domain.Product{Name: dto.Name, CurrentPrice: dto.CurrentPrice}
}

type StoreInventoryResponse struct {
	StoreID       int `json:"store_id"`
	ProductID     int `json:"product_id"`
	StockQuantity int `json:"stock_quantity"`
}

func ToStoreInventoryResponse(i domain.StoreInventory) StoreInventoryResponse {
	return StoreInventoryResponse{StoreID: i.StoreID, ProductID: i.ProductID, StockQuantity: i.StockQuantity}
}
