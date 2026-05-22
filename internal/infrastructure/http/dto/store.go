package dto

import "trade-system/internal/domain"

// StoreResponse — структура для отправки JSON на клиент
type StoreResponse struct {
	ID           int     `json:"id"`
	Area         float64 `json:"area"`
	CounterCount int     `json:"counter_count"`
	Type         string  `json:"store_type"`
}

// ToStoreResponse маппит доменную сущность в JSON-ответ
func ToStoreResponse(s domain.Store) StoreResponse {
	return StoreResponse{
		ID:           s.ID,
		Area:         s.Area,
		CounterCount: s.CounterCount,
		Type:         s.Type,
	}
}

// CreateStoreRequest — структура для приема данных от клиента (POST запрос)
type CreateStoreRequest struct {
	Area         float64 `json:"area" binding:"required,gt=0"` // Валидация Gin
	CounterCount int     `json:"counter_count" binding:"required,gte=0"`
	Type         string  `json:"store_type" binding:"required"`
}

// ToDomain конвертирует запрос в доменную сущность (ID будет 0, т.к. точка еще не создана)
func (req *CreateStoreRequest) ToDomain() domain.Store {
	return domain.Store{
		Area:         req.Area,
		CounterCount: req.CounterCount,
		Type:         req.Type,
	}
}
