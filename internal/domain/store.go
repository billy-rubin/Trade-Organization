package domain

type Store struct {
	ID           int
	Area         float64
	CounterCount int
	Type         string
}

type Section struct {
	ID      int
	Name    string
	Floor   int
	StoreID int
}

type Hall struct {
	ID      int
	Name    string
	StoreID int
}
