package domain

type Product struct {
	ID           int
	Name         string
	CurrentPrice float64
}

type StoreInventory struct {
	StoreID       int
	ProductID     int
	StockQuantity int
}

type SupplierCatalog struct {
	SupplierID  int
	ProductID   int
	SupplyPrice float64
}
