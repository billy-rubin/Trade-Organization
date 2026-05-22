package dto

import "trade-organization/internal/domain"

type StoreResponse struct {
	ID           int     `json:"id"`
	Area         float64 `json:"area"`
	CounterCount int     `json:"counter_count"`
	Type         string  `json:"store_type"`
}

func ToStoreResponse(s domain.Store) StoreResponse {
	return StoreResponse{
		ID:           s.ID,
		Area:         s.Area,
		CounterCount: s.CounterCount,
		Type:         s.Type,
	}
}

type CreateStoreRequest struct {
	Area         float64 `json:"area" binding:"required,gt=0"` // Валидация Gin
	CounterCount int     `json:"counter_count" binding:"required,gte=0"`
	Type         string  `json:"store_type" binding:"required"`
}

func (req *CreateStoreRequest) ToDomain() domain.Store {
	return domain.Store{
		Area:         req.Area,
		CounterCount: req.CounterCount,
		Type:         req.Type,
	}
}
