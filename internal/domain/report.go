package domain

import "time"

type InventoryReport struct {
	StoreID       int
	StoreType     string
	ProductID     int
	ProductName   string
	StockQuantity int
}

type SellerOutputReport struct {
	SellerID       int
	SellerName     string
	StoreID        int
	TotalRevenue   float64
	TotalItemsSold int
}

type ProfitabilityReport struct {
	StoreID            int
	TotalRevenue       float64
	TotalOverhead      float64
	ProfitabilityRatio *float64
}

type ActiveCustomerReport struct {
	CustomerID    int
	CustomerName  string
	TotalReceipts int
	TotalSpent    float64
}

type TurnoverReport struct {
	StoreType      string
	TotalItemsSold int
	TotalTurnover  float64
}

type SupplierDeliveryReport struct {
	OrderID      int       `json:"order_id"`
	OrderDate    time.Time `json:"order_date"`
	SupplierName string    `json:"supplier_name"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	SupplyPrice  float64   `json:"supply_price"`
	TotalCost    float64   `json:"total_cost"`
}

type StoreEfficiencyReport struct {
	StoreID           int     `json:"store_id"`
	StoreType         string  `json:"store_type"`
	Area              float64 `json:"area"`
	CounterCount      int     `json:"counter_count"`
	TotalRevenue      float64 `json:"total_revenue"`
	RevenuePerSqMeter float64 `json:"revenue_per_sq_meter"`
}

type ProductCustomerReport struct {
	CustomerID      int     `json:"customer_id"`
	CustomerName    string  `json:"customer_name"`
	Characteristics *string `json:"characteristics"`
	BoughtProduct   string  `json:"bought_product"`
}
