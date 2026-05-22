package dto

import "trade-organization/internal/domain"

type InventoryReportResponse struct {
	StoreID       int    `json:"store_id"`
	StoreType     string `json:"store_type"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	StockQuantity int    `json:"stock_quantity"`
}

func ToInventoryReportResponseList(domainList []domain.InventoryReport) []InventoryReportResponse {
	result := make([]InventoryReportResponse, 0, len(domainList))
	for _, item := range domainList {
		result = append(result, InventoryReportResponse{
			StoreID:       item.StoreID,
			StoreType:     item.StoreType,
			ProductID:     item.ProductID,
			ProductName:   item.ProductName,
			StockQuantity: item.StockQuantity,
		})
	}
	return result
}

type ProfitabilityResponse struct {
	StoreID            int      `json:"store_id"`
	TotalRevenue       float64  `json:"total_revenue"`
	TotalOverhead      float64  `json:"total_overhead"`
	ProfitabilityRatio *float64 `json:"profitability_ratio"`
}

func ToProfitabilityResponse(p domain.ProfitabilityReport) ProfitabilityResponse {
	return ProfitabilityResponse{
		StoreID:            p.StoreID,
		TotalRevenue:       p.TotalRevenue,
		TotalOverhead:      p.TotalOverhead,
		ProfitabilityRatio: p.ProfitabilityRatio,
	}
}

type TurnoverResponse struct {
	StoreType      string  `json:"store_type"`
	TotalItemsSold int     `json:"total_items_sold"`
	TotalTurnover  float64 `json:"total_turnover"`
}

func ToTurnoverResponseList(domainList []domain.TurnoverReport) []TurnoverResponse {
	result := make([]TurnoverResponse, 0, len(domainList))
	for _, item := range domainList {
		result = append(result, TurnoverResponse{
			StoreType:      item.StoreType,
			TotalItemsSold: item.TotalItemsSold,
			TotalTurnover:  item.TotalTurnover,
		})
	}
	return result
}

type SellerOutputResponse struct {
	SellerID       int     `json:"seller_id"`
	SellerName     string  `json:"seller_name"`
	StoreID        int     `json:"store_id"`
	TotalRevenue   float64 `json:"total_revenue"`
	TotalItemsSold int     `json:"total_items_sold"`
}

func ToSellerOutputResponse(s domain.SellerOutputReport) SellerOutputResponse {
	return SellerOutputResponse{
		SellerID:       s.SellerID,
		SellerName:     s.SellerName,
		StoreID:        s.StoreID,
		TotalRevenue:   s.TotalRevenue,
		TotalItemsSold: s.TotalItemsSold,
	}
}

type ActiveCustomerResponse struct {
	CustomerID    int     `json:"customer_id"`
	CustomerName  string  `json:"customer_name"`
	TotalReceipts int     `json:"total_receipts"`
	TotalSpent    float64 `json:"total_spent"`
}

func ToActiveCustomerResponseList(domainList []domain.ActiveCustomerReport) []ActiveCustomerResponse {
	result := make([]ActiveCustomerResponse, 0, len(domainList))
	for _, item := range domainList {
		result = append(result, ActiveCustomerResponse{
			CustomerID:    item.CustomerID,
			CustomerName:  item.CustomerName,
			TotalReceipts: item.TotalReceipts,
			TotalSpent:    item.TotalSpent,
		})
	}
	return result
}
