CREATE OR REPLACE FUNCTION tf_ValidateCustomerStoreType()
RETURNS TRIGGER AS $$
DECLARE
v_store_type VARCHAR(50);
BEGIN
SELECT s.store_type INTO v_store_type
FROM Seller sel
         JOIN Store s ON sel.store_id = s.store_id
WHERE sel.seller_id = NEW.seller_id;

IF NEW.customer_id IS NOT NULL AND v_store_type IN ('Kiosk', 'Tray') THEN
        RAISE EXCEPTION 'Compliance Error: Customers cannot be registered for Kiosks or Trays.';
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_ValidateCustomerStoreType
    BEFORE INSERT ON Sale
    FOR EACH ROW
    EXECUTE FUNCTION tf_ValidateCustomerStoreType();

-- 1. Функция и триггер защиты от отрицательных остатков
CREATE OR REPLACE FUNCTION tf_PreventNegativeStock()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.stock_quantity < 0 THEN
        RAISE EXCEPTION 'Inventory Error: Stock quantity cannot be negative (Product ID: %, Store ID: %).', NEW.product_id, NEW.store_id;
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_PreventNegativeStock
    BEFORE UPDATE OR INSERT ON Store_Inventory
    FOR EACH ROW
EXECUTE FUNCTION tf_PreventNegativeStock();


-- 2. Функция и триггер автосписания товаров при продаже
CREATE OR REPLACE FUNCTION tf_AutoDeductItem()
RETURNS TRIGGER AS $$
DECLARE
v_store_id INT;
    v_current_stock INT;
BEGIN
    -- Находим магазин, в котором пробит чек, через связь с таблицей Sale и Seller
SELECT sel.store_id INTO v_store_id
FROM Sale s
         JOIN Seller sel ON s.seller_id = sel.seller_id
WHERE s.sale_id = NEW.sale_id;

-- Получаем текущий остаток товара в этом магазине
SELECT stock_quantity INTO v_current_stock
FROM Store_Inventory
WHERE store_id = v_store_id AND product_id = NEW.product_id;

-- Проверяем, хватает ли товара
IF v_current_stock IS NULL OR v_current_stock < NEW.quantity THEN
        RAISE EXCEPTION 'Sale Error: Not enough stock for Product ID: % in Store ID: %', NEW.product_id, v_store_id;
END IF;

    -- Уменьшаем количество на складе
UPDATE Store_Inventory
SET stock_quantity = stock_quantity - NEW.quantity
WHERE store_id = v_store_id AND product_id = NEW.product_id;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_AutoDeductItem
    AFTER INSERT ON Sale_Details
    FOR EACH ROW
    EXECUTE FUNCTION tf_AutoDeductItem();

-- Хранимые процедуры
CREATE OR REPLACE PROCEDURE sp_GenerateOrderFromRequest(
    p_request_id INT,
    p_supplier_id INT
)
LANGUAGE plpgsql
AS $$
DECLARE
v_new_order_id INT;
    v_check_request INT;
BEGIN
SELECT request_id INTO v_check_request
FROM Request
WHERE request_id = p_request_id AND status != 'Processed'
FOR UPDATE;

IF NOT FOUND THEN
        RAISE EXCEPTION 'Request not found or already processed.';
END IF;

INSERT INTO Purchase_Order (order_date, status, supplier_id)
VALUES (CURRENT_DATE, 'Created', p_supplier_id)
    RETURNING order_id INTO v_new_order_id;

INSERT INTO Order_Details (order_id, product_id, quantity)
SELECT v_new_order_id, product_id, quantity
FROM Request_Details
WHERE request_id = p_request_id;

UPDATE Request
SET status = 'Processed'
WHERE request_id = p_request_id;

COMMIT;
END;
$$;

CREATE OR REPLACE PROCEDURE sp_TransferProduct(
    p_from_store_id INT,
    p_to_store_id INT,
    p_product_id INT,
    p_quantity INT
)
LANGUAGE plpgsql
AS $$
DECLARE
v_current_stock INT;
BEGIN
SELECT stock_quantity INTO v_current_stock
FROM Store_Inventory
WHERE store_id = p_from_store_id AND product_id = p_product_id;

IF v_current_stock IS NULL OR v_current_stock < p_quantity THEN
        RAISE EXCEPTION 'Transfer failed: Insufficient stock in source store.';
END IF;

UPDATE Store_Inventory
SET stock_quantity = stock_quantity - p_quantity
WHERE store_id = p_from_store_id AND product_id = p_product_id;

INSERT INTO Store_Inventory (store_id, product_id, stock_quantity)
VALUES (p_to_store_id, p_product_id, p_quantity)
    ON CONFLICT (store_id, product_id) DO UPDATE
                                              SET stock_quantity = Store_Inventory.stock_quantity + p_quantity;

COMMIT;
END;
$$;

CREATE OR REPLACE PROCEDURE sp_ReceiveOrder(
    p_order_id INT,
    p_store_id INT
)
LANGUAGE plpgsql
AS $$
DECLARE
r_item RECORD;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM Purchase_Order WHERE order_id = p_order_id AND status = 'Created') THEN
        RAISE EXCEPTION 'Order cannot be received: Invalid status or order not found.';
END IF;

FOR r_item IN SELECT product_id, quantity FROM Order_Details WHERE order_id = p_order_id LOOP
              INSERT INTO Store_Inventory (store_id, product_id, stock_quantity)
              VALUES (p_store_id, r_item.product_id, r_item.quantity)
              ON CONFLICT (store_id, product_id) DO UPDATE
                                                        SET stock_quantity = Store_Inventory.stock_quantity + r_item.quantity;
END LOOP;

UPDATE Purchase_Order
SET status = 'Delivered'
WHERE order_id = p_order_id;

COMMIT;
END;
$$;

-- Представления (Views)
CREATE OR REPLACE VIEW v_StoreInventory AS
SELECT
    s.store_id,
    s.store_type,
    p.product_id,
    p.name AS product_name,
    si.stock_quantity
FROM Store s
         JOIN Store_Inventory si ON s.store_id = si.store_id
         JOIN Product p ON si.product_id = p.product_id;

CREATE OR REPLACE VIEW v_ActiveCustomers AS
SELECT
    c.customer_id,
    c.full_name AS customer_name,
    COUNT(DISTINCT s.sale_id) AS total_receipts,
    SUM(sd.quantity * p.current_price) AS total_spent
FROM Customer c
         JOIN Sale s ON c.customer_id = s.customer_id
         JOIN Sale_Details sd ON s.sale_id = sd.sale_id
         JOIN Product p ON sd.product_id = p.product_id
GROUP BY c.customer_id, c.full_name;

CREATE OR REPLACE VIEW v_TurnoverByStoreType AS
SELECT
    st.store_type,
    SUM(sd.quantity) AS total_items_sold,
    SUM(sd.quantity * p.current_price) AS total_turnover
FROM Store st
         JOIN Seller sel ON st.store_id = sel.store_id
         JOIN Sale s ON sel.seller_id = s.seller_id
         JOIN Sale_Details sd ON s.sale_id = sd.sale_id
         JOIN Product p ON sd.product_id = p.product_id
GROUP BY st.store_type;

-- Роли и Пользователи
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'role_seller') THEN
CREATE ROLE role_seller;
END IF;
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'role_purchase_manager') THEN
CREATE ROLE role_purchase_manager;
END IF;
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'role_director') THEN
CREATE ROLE role_director;
END IF;
END
$$;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO role_seller, role_purchase_manager, role_director;
GRANT SELECT ON Store, Product, Store_Inventory, Request, Request_Details, Customer, Sale, Sale_Details TO role_seller;
GRANT INSERT ON Sale, Sale_Details, Request, Request_Details, Customer TO role_seller;
GRANT UPDATE ON Customer, Product, Store_Inventory TO role_seller;

GRANT SELECT ON Store, Request, Request_Details, Product, Store_Inventory, Purchase_Order, Order_Details, Supplier, Supplier_Catalog TO role_purchase_manager;
GRANT INSERT ON Purchase_Order, Order_Details, Supplier, Supplier_Catalog, Product TO role_purchase_manager;
GRANT UPDATE ON Purchase_Order, Order_Details, Supplier, Supplier_Catalog, Product, Request, Store_Inventory TO role_purchase_manager;
GRANT DELETE ON Supplier, Supplier_Catalog TO role_purchase_manager;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO role_director;
GRANT INSERT, UPDATE ON Store, Department_Store, Shop, Kiosk, Tray, Section, Hall, Seller, Position, Payment TO role_director;
GRANT DELETE ON Seller TO role_director;