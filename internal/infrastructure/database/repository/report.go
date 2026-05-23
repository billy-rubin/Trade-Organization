package repository

import (
	"context"
	"fmt"
	"time"

	"trade-organization/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportRepository struct {
	pool *pgxpool.Pool
}

func NewReportRepository(pool *pgxpool.Pool) *ReportRepository {
	return &ReportRepository{pool: pool}
}

func (r *ReportRepository) GetStoreInventory(ctx context.Context, storeType string) ([]domain.InventoryReport, error) {
	query := `
		SELECT store_id, store_type, product_id, product_name, stock_quantity 
		FROM v_StoreInventory 
		WHERE store_type = $1 
		ORDER BY store_id, product_name;
	`
	rows, err := r.pool.Query(ctx, query, storeType)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []domain.InventoryReport
	for rows.Next() {
		var item domain.InventoryReport
		if err := rows.Scan(&item.StoreID, &item.StoreType, &item.ProductID, &item.ProductName, &item.StockQuantity); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, item)
	}
	return results, nil
}

func (r *ReportRepository) GetSellerOutput(ctx context.Context, sellerID int, startDate, endDate time.Time) (domain.SellerOutputReport, error) {
	query := `
		SELECT 
			sel.seller_id, sel.full_name AS seller_name, st.store_id,
			COALESCE(SUM(sd.quantity * p.current_price), 0) AS total_revenue,
			COALESCE(SUM(sd.quantity), 0) AS total_items_sold
		FROM Seller sel
		JOIN Store st ON sel.store_id = st.store_id
		JOIN Sale s ON sel.seller_id = s.seller_id
		JOIN Sale_Details sd ON s.sale_id = sd.sale_id
		JOIN Product p ON sd.product_id = p.product_id
		WHERE sel.seller_id = $1 AND s.sale_date >= $2 AND s.sale_date <= $3
		GROUP BY sel.seller_id, sel.full_name, st.store_id;
	`
	var report domain.SellerOutputReport
	err := r.pool.QueryRow(ctx, query, sellerID, startDate, endDate).Scan(
		&report.SellerID, &report.SellerName, &report.StoreID, &report.TotalRevenue, &report.TotalItemsSold,
	)
	if err != nil {
		return domain.SellerOutputReport{}, fmt.Errorf("failed to get seller output: %w", err)
	}
	return report, nil
}

func (r *ReportRepository) GetStoreProfitability(ctx context.Context, storeID int, startDate, endDate time.Time) (domain.ProfitabilityReport, error) {
	query := `
		WITH SalesRevenue AS (
			SELECT sel.store_id, SUM(sd.quantity * p.current_price) AS revenue
			FROM Sale s
			JOIN Seller sel ON s.seller_id = sel.seller_id
			JOIN Sale_Details sd ON s.sale_id = sd.sale_id
			JOIN Product p ON sd.product_id = p.product_id
			WHERE s.sale_date >= $2 AND s.sale_date <= $3
			GROUP BY sel.store_id
		),
		OverheadPayments AS (
			SELECT store_id, SUM(amount) AS total_payments
			FROM Payment
			WHERE payment_date >= $2 AND payment_date <= $3
			GROUP BY store_id
		),
		OverheadSalaries AS (
			SELECT sel.store_id, SUM(pos.salary) AS total_salaries
			FROM Seller sel
			JOIN Position pos ON sel.position_id = pos.position_id
			GROUP BY sel.store_id
		)
		SELECT 
			st.store_id,
			COALESCE(sr.revenue, 0) AS total_revenue,
			(COALESCE(op.total_payments, 0) + COALESCE(os.total_salaries, 0)) AS total_overhead,
			CASE 
				WHEN (COALESCE(op.total_payments, 0) + COALESCE(os.total_salaries, 0)) = 0 THEN NULL
				ELSE COALESCE(sr.revenue, 0) / (COALESCE(op.total_payments, 0) + COALESCE(os.total_salaries, 0))
			END AS profitability_ratio
		FROM Store st
		LEFT JOIN SalesRevenue sr ON st.store_id = sr.store_id
		LEFT JOIN OverheadPayments op ON st.store_id = op.store_id
		LEFT JOIN OverheadSalaries os ON st.store_id = os.store_id
		WHERE st.store_id = $1;
	`
	var report domain.ProfitabilityReport
	err := r.pool.QueryRow(ctx, query, storeID, startDate, endDate).Scan(
		&report.StoreID, &report.TotalRevenue, &report.TotalOverhead, &report.ProfitabilityRatio,
	)
	if err != nil {
		return domain.ProfitabilityReport{}, fmt.Errorf("failed to calculate profitability: %w", err)
	}
	return report, nil
}

func (r *ReportRepository) GetTopActiveCustomers(ctx context.Context, limit int) ([]domain.ActiveCustomerReport, error) {
	query := `
		SELECT customer_id, customer_name, total_receipts, total_spent 
		FROM v_ActiveCustomers 
		ORDER BY total_spent DESC 
		FETCH FIRST $1 ROWS ONLY
	`
	rows, err := r.pool.Query(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []domain.ActiveCustomerReport
	for rows.Next() {
		var item domain.ActiveCustomerReport
		if err := rows.Scan(&item.CustomerID, &item.CustomerName, &item.TotalReceipts, &item.TotalSpent); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, item)
	}
	return results, nil
}

func (r *ReportRepository) GetTurnoverByStoreType(ctx context.Context) ([]domain.TurnoverReport, error) {
	query := `
		SELECT store_type, total_items_sold, total_turnover 
		FROM v_TurnoverByStoreType;
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var results []domain.TurnoverReport
	for rows.Next() {
		var item domain.TurnoverReport
		if err := rows.Scan(&item.StoreType, &item.TotalItemsSold, &item.TotalTurnover); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, item)
	}
	return results, nil
}

func (r *ReportRepository) GetSupplierDeliveries(ctx context.Context, supplierID, productID int, startDate, endDate time.Time) ([]domain.SupplierDeliveryReport, error) {
	query := `
		SELECT 
			po.order_id, po.order_date, sup.name AS supplier_name, 
			p.name AS product_name, od.quantity, sc.supply_price, 
			(od.quantity * sc.supply_price) AS total_cost
		FROM Purchase_Order po
		JOIN Supplier sup ON po.supplier_id = sup.supplier_id
		JOIN Order_Details od ON po.order_id = od.order_id
		JOIN Product p ON od.product_id = p.product_id
		JOIN Supplier_Catalog sc ON sup.supplier_id = sc.supplier_id AND p.product_id = sc.product_id
		WHERE sup.supplier_id = $1 AND p.product_id = $2 AND po.status = 'Delivered'
		  AND po.order_date BETWEEN $3 AND $4
		ORDER BY po.order_date DESC;
	`

	rows, err := r.pool.Query(ctx, query, supplierID, productID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.SupplierDeliveryReport
	for rows.Next() {
		var item domain.SupplierDeliveryReport
		err := rows.Scan(
			&item.OrderID, &item.OrderDate, &item.SupplierName,
			&item.ProductName, &item.Quantity, &item.SupplyPrice, &item.TotalCost,
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

func (r *ReportRepository) GetStoreEfficiency(ctx context.Context, storeID int) (domain.StoreEfficiencyReport, error) {
	query := `
		SELECT 
			st.store_id, st.store_type, st.area, st.counter_count,
			COALESCE(SUM(sd.quantity * p.current_price), 0) AS total_revenue,
			COALESCE(SUM(sd.quantity * p.current_price) / NULLIF(st.area, 0), 0) AS revenue_per_sq_meter
		FROM Store st
		LEFT JOIN Seller sel ON st.store_id = sel.store_id
		LEFT JOIN Sale s ON sel.seller_id = s.seller_id
		LEFT JOIN Sale_Details sd ON s.sale_id = sd.sale_id
		LEFT JOIN Product p ON sd.product_id = p.product_id
		WHERE st.store_id = $1
		GROUP BY st.store_id, st.store_type, st.area, st.counter_count;
	`

	var item domain.StoreEfficiencyReport
	err := r.pool.QueryRow(ctx, query, storeID).Scan(
		&item.StoreID, &item.StoreType, &item.Area, &item.CounterCount,
		&item.TotalRevenue, &item.RevenuePerSqMeter,
	)
	if err != nil {
		return item, err // pgx.ErrNoRows обработается на уровне выше
	}

	return item, nil
}

func (r *ReportRepository) GetProductCustomers(ctx context.Context, productID, storeID int) ([]domain.ProductCustomerReport, error) {
	query := `
		SELECT DISTINCT 
			c.customer_id, c.full_name AS customer_name, c.characteristics, p.name AS bought_product
		FROM Customer c
		JOIN Sale s ON c.customer_id = s.customer_id
		JOIN Sale_Details sd ON s.sale_id = sd.sale_id
		JOIN Product p ON sd.product_id = p.product_id
		JOIN Seller sel ON s.seller_id = sel.seller_id
		WHERE p.product_id = $1 AND sel.store_id = $2;
	`

	rows, err := r.pool.Query(ctx, query, productID, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.ProductCustomerReport
	for rows.Next() {
		var item domain.ProductCustomerReport
		err := rows.Scan(
			&item.CustomerID, &item.CustomerName, &item.Characteristics, &item.BoughtProduct,
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
