package database

import (
	"time"
	"trade-organization/internal/domain"
)

type StoreDB struct {
	ID           int     `db:"store_id"`
	Area         float64 `db:"area"`
	CounterCount int     `db:"counter_count"`
	Type         string  `db:"store_type"`
}

func (s *StoreDB) ToDomain() domain.Store {
	return domain.Store{ID: s.ID, Area: s.Area, CounterCount: s.CounterCount, Type: s.Type}
}

type SectionDB struct {
	ID      int    `db:"section_id"`
	Name    string `db:"name"`
	Floor   int    `db:"floor"`
	StoreID int    `db:"store_id"`
}

func (s *SectionDB) ToDomain() domain.Section {
	return domain.Section{ID: s.ID, Name: s.Name, Floor: s.Floor, StoreID: s.StoreID}
}

type HallDB struct {
	ID      int    `db:"hall_id"`
	Name    string `db:"name"`
	StoreID int    `db:"store_id"`
}

func (h *HallDB) ToDomain() domain.Hall {
	return domain.Hall{ID: h.ID, Name: h.Name, StoreID: h.StoreID}
}

type PositionDB struct {
	ID     int     `db:"position_id"`
	Title  string  `db:"title"`
	Salary float64 `db:"salary"`
}

func (p *PositionDB) ToDomain() domain.Position {
	return domain.Position{ID: p.ID, Title: p.Title, Salary: p.Salary}
}

type SellerDB struct {
	ID         int    `db:"seller_id"`
	FullName   string `db:"full_name"`
	PositionID int    `db:"position_id"`
	StoreID    int    `db:"store_id"`
}

func (s *SellerDB) ToDomain() domain.Seller {
	return domain.Seller{ID: s.ID, FullName: s.FullName, PositionID: s.PositionID, StoreID: s.StoreID}
}

type UserDB struct {
	ID           int    `db:"user_id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role"`
	SellerID     *int   `db:"seller_id"`
}

func (u *UserDB) ToDomain() domain.User {
	return domain.User{ID: u.ID, Username: u.Username, PasswordHash: u.PasswordHash, Role: u.Role, SellerID: u.SellerID}
}

type ProductDB struct {
	ID           int     `db:"product_id"`
	Name         string  `db:"name"`
	CurrentPrice float64 `db:"current_price"`
}

func (p *ProductDB) ToDomain() domain.Product {
	return domain.Product{ID: p.ID, Name: p.Name, CurrentPrice: p.CurrentPrice}
}

type StoreInventoryDB struct {
	StoreID       int `db:"store_id"`
	ProductID     int `db:"product_id"`
	StockQuantity int `db:"stock_quantity"`
}

func (i *StoreInventoryDB) ToDomain() domain.StoreInventory {
	return domain.StoreInventory{StoreID: i.StoreID, ProductID: i.ProductID, StockQuantity: i.StockQuantity}
}

type SupplierCatalogDB struct {
	SupplierID  int     `db:"supplier_id"`
	ProductID   int     `db:"product_id"`
	SupplyPrice float64 `db:"supply_price"`
}

func (c *SupplierCatalogDB) ToDomain() domain.SupplierCatalog {
	return domain.SupplierCatalog{SupplierID: c.SupplierID, ProductID: c.ProductID, SupplyPrice: c.SupplyPrice}
}

type CustomerDB struct {
	ID              int    `db:"customer_id"`
	FullName        string `db:"full_name"`
	Characteristics string `db:"characteristics"`
}

func (c *CustomerDB) ToDomain() domain.Customer {
	return domain.Customer{ID: c.ID, FullName: c.FullName, Characteristics: c.Characteristics}
}

type SaleDB struct {
	ID         int       `db:"sale_id"`
	Date       time.Time `db:"sale_date"`
	SellerID   int       `db:"seller_id"`
	CustomerID *int      `db:"customer_id"`
}

func (s *SaleDB) ToDomain() domain.Sale {
	return domain.Sale{ID: s.ID, Date: s.Date, SellerID: s.SellerID, CustomerID: s.CustomerID}
}

type SupplierDB struct {
	ID   int    `db:"supplier_id"`
	Name string `db:"name"`
}

func (s *SupplierDB) ToDomain() domain.Supplier {
	return domain.Supplier{ID: s.ID, Name: s.Name}
}

type RequestDB struct {
	ID      int       `db:"request_id"`
	Date    time.Time `db:"request_date"`
	Status  string    `db:"status"`
	StoreID int       `db:"store_id"`
}

func (r *RequestDB) ToDomain() domain.Request {
	return domain.Request{ID: r.ID, Date: r.Date, Status: r.Status, StoreID: r.StoreID}
}

type PurchaseOrderDB struct {
	ID         int       `db:"order_id"`
	Date       time.Time `db:"order_date"`
	Status     string    `db:"status"`
	SupplierID int       `db:"supplier_id"`
}

func (po *PurchaseOrderDB) ToDomain() domain.PurchaseOrder {
	return domain.PurchaseOrder{ID: po.ID, Date: po.Date, Status: po.Status, SupplierID: po.SupplierID}
}

type PaymentDB struct {
	ID      int       `db:"payment_id"`
	Type    string    `db:"payment_type"`
	Amount  float64   `db:"amount"`
	Date    time.Time `db:"payment_date"`
	StoreID int       `db:"store_id"`
}

func (p *PaymentDB) ToDomain() domain.Payment {
	return domain.Payment{ID: p.ID, Type: p.Type, Amount: p.Amount, Date: p.Date, StoreID: p.StoreID}
}
