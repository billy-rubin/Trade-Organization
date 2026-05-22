package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"trade-organization/internal/domain"
)

type SupplyRepository struct {
	pool *pgxpool.Pool
}

func NewSupplyRepository(pool *pgxpool.Pool) *SupplyRepository {
	return &SupplyRepository{pool: pool}
}

func (r *SupplyRepository) CreateRequest(ctx context.Context, req domain.Request, details []domain.RequestDetail) (int, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var requestID int
	queryReq := `
		INSERT INTO Request (request_date, status, store_id)
		VALUES ($1, $2, $3)
		RETURNING request_id
	`

	err = tx.QueryRow(ctx, queryReq, req.Date, req.Status, req.StoreID).Scan(&requestID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert request: %w", err)
	}

	batch := &pgx.Batch{}
	queryDetail := `
		INSERT INTO Request_Details (request_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`
	for _, item := range details {
		batch.Queue(queryDetail, requestID, item.ProductID, item.Quantity)
	}

	br := tx.SendBatch(ctx, batch)
	for i := 0; i < len(details); i++ {
		if _, err := br.Exec(); err != nil {
			br.Close()
			return 0, fmt.Errorf("failed to insert request detail row %d: %w", i, err)
		}
	}
	br.Close()

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return requestID, nil
}

// GenerateOrderFromRequest вызывает хранимую процедуру для автоматической генерации заказа поставщику на основе заявки
func (r *SupplyRepository) GenerateOrderFromRequest(ctx context.Context, requestID int, supplierID int) error {
	query := `CALL sp_GenerateOrderFromRequest($1, $2)`

	_, err := r.pool.Exec(ctx, query, requestID, supplierID)
	if err != nil {
		return fmt.Errorf("failed to generate order from request via procedure: %w", err)
	}

	return nil
}

// ReceiveOrder вызывает хранимую процедуру для оприходования товаров из заказа на склад торговой точки
func (r *SupplyRepository) ReceiveOrder(ctx context.Context, orderID int, storeID int) error {
	query := `CALL sp_ReceiveOrder($1, $2)`

	_, err := r.pool.Exec(ctx, query, orderID, storeID)
	if err != nil {
		return fmt.Errorf("failed to receive order via procedure: %w", err)
	}

	return nil
}

// Запрос 9: Получить номенклатуру и объем товаров в указанном заказе
func (r *SupplyRepository) GetOrderDetails(ctx context.Context, orderID int) ([]domain.OrderDetailReport, error) {
	query := `
		SELECT 
			po.order_id, po.order_date, po.status, 
			sup.name AS supplier_name, p.name AS product_name, od.quantity
		FROM Purchase_Order po
		JOIN Supplier sup ON po.supplier_id = sup.supplier_id
		JOIN Order_Details od ON po.order_id = od.order_id
		JOIN Product p ON od.product_id = p.product_id
		WHERE po.order_id = $1;
	`

	rows, err := r.pool.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.OrderDetailReport
	for rows.Next() {
		var item domain.OrderDetailReport
		err := rows.Scan(
			&item.OrderID, &item.OrderDate, &item.Status,
			&item.SupplierName, &item.ProductName, &item.Quantity,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
