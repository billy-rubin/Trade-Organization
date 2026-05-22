package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"trade-organization/internal/domain"
)

type TradeRepository struct {
	pool *pgxpool.Pool
}

func NewTradeRepository(pool *pgxpool.Pool) *TradeRepository {
	return &TradeRepository{pool: pool}
}

func (r *TradeRepository) CreateSale(ctx context.Context, sale domain.Sale, details []domain.SaleDetail) (int, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var saleID int
	querySale := `
		INSERT INTO Sale (sale_date, seller_id, customer_id)
		VALUES ($1, $2, $3)
		RETURNING sale_id
	`

	err = tx.QueryRow(ctx, querySale, sale.Date, sale.SellerID, sale.CustomerID).Scan(&saleID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert sale: %w", err)
	}

	batch := &pgx.Batch{}
	queryDetail := `
		INSERT INTO Sale_Details (sale_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`
	for _, item := range details {
		batch.Queue(queryDetail, saleID, item.ProductID, item.Quantity)
	}

	br := tx.SendBatch(ctx, batch)
	for i := 0; i < len(details); i++ {
		if _, err := br.Exec(); err != nil {
			br.Close()
			return 0, fmt.Errorf("failed to insert sale detail row %d: %w", i, err)
		}
	}
	br.Close()

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return saleID, nil
}

func (r *TradeRepository) TransferProduct(ctx context.Context, fromStoreID, toStoreID, productID, quantity int) error {
	query := `CALL sp_TransferProduct($1, $2, $3, $4)`

	_, err := r.pool.Exec(ctx, query, fromStoreID, toStoreID, productID, quantity)
	if err != nil {
		return fmt.Errorf("database error during product transfer: %w", err)
	}

	return nil
}
