package domain

import "time"

type Customer struct {
	ID              int
	FullName        string
	Characteristics string
}

type Sale struct {
	ID         int
	Date       time.Time
	SellerID   int
	CustomerID *int
}

type SaleDetail struct {
	SaleID    int
	ProductID int
	Quantity  int
}
