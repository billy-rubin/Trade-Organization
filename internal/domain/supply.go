package domain

import "time"

type Supplier struct {
	ID   int
	Name string
}

type Request struct {
	ID      int
	Date    time.Time
	Status  string
	StoreID int
}

type RequestDetail struct {
	RequestID int
	ProductID int
	Quantity  int
}

type PurchaseOrder struct {
	ID         int
	Date       time.Time
	Status     string
	SupplierID int
}

type OrderDetail struct {
	OrderID   int
	ProductID int
	Quantity  int
}

type Payment struct {
	ID      int
	Type    string
	Amount  float64
	Date    time.Time
	StoreID int
}

type OrderDetailReport struct {
	OrderID      int       `json:"order_id"`
	OrderDate    time.Time `json:"order_date"`
	Status       string    `json:"status"`
	SupplierName string    `json:"supplier_name"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
}
