package dto

import "trade-organization/internal/domain"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type SellerResponse struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	PositionID int    `json:"position_id"`
	StoreID    int    `json:"store_id"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
	SellerID *int   `json:"seller_id"`
}

func ToSellerResponse(s domain.Seller) SellerResponse {
	return SellerResponse{ID: s.ID, FullName: s.FullName, PositionID: s.PositionID, StoreID: s.StoreID}
}
