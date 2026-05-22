package domain

type Position struct {
	ID     int
	Title  string
	Salary float64
}

type Seller struct {
	ID         int
	FullName   string
	PositionID int
	StoreID    int
}

type User struct {
	ID           int
	Username     string
	PasswordHash string
	Role         string
	SellerID     *int
}
