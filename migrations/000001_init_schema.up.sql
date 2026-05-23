CREATE TABLE IF NOT EXIST Store (
                       store_id SERIAL PRIMARY KEY,
                       area NUMERIC(10, 2) NOT NULL CHECK (area > 0),
                       counter_count INT NOT NULL CHECK (counter_count >= 0),
                       store_type VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXIST Position (
                          position_id SERIAL PRIMARY KEY,
                          title VARCHAR(100) UNIQUE NOT NULL,
                          salary NUMERIC(10, 2) NOT NULL CHECK (salary >= 0)
);

CREATE TABLE IF NOT EXIST Customer (
                          customer_id SERIAL PRIMARY KEY,
                          full_name VARCHAR(150) NOT NULL,
                          characteristics TEXT
);

CREATE TABLE IF NOT EXIST Product (
                         product_id SERIAL PRIMARY KEY,
                         name VARCHAR(150) NOT NULL,
                         current_price NUMERIC(10, 2) NOT NULL CHECK (current_price >= 0)
);

CREATE TABLE IF NOT EXIST Supplier (
                          supplier_id SERIAL PRIMARY KEY,
                          name VARCHAR(150) NOT NULL
);

CREATE TABLE IF NOT EXIST Department_Store (
                                  store_id INT PRIMARY KEY REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Shop (
                      store_id INT PRIMARY KEY REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Kiosk (
                       store_id INT PRIMARY KEY REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Tray (
                      store_id INT PRIMARY KEY REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Section (
                         section_id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         floor INT,
                         store_id INT NOT NULL REFERENCES Department_Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Hall (
                      hall_id SERIAL PRIMARY KEY,
                      name VARCHAR(100) NOT NULL,
                      store_id INT NOT NULL REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Seller (
                        seller_id SERIAL PRIMARY KEY,
                        full_name VARCHAR(150) NOT NULL,
                        position_id INT NOT NULL REFERENCES Position(position_id) ON DELETE RESTRICT,
                        store_id INT NOT NULL REFERENCES Store(store_id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXIST App_Users (
                           user_id SERIAL PRIMARY KEY,
                           username VARCHAR(50) UNIQUE NOT NULL,
                           password_hash VARCHAR(255) NOT NULL,
                           role VARCHAR(50) NOT NULL,
                           seller_id INT REFERENCES Seller(seller_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXIST Payment (
                         payment_id SERIAL PRIMARY KEY,
                         payment_type VARCHAR(50) NOT NULL,
                         amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
                         payment_date DATE NOT NULL,
                         store_id INT NOT NULL REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Store_Inventory (
                                 store_id INT NOT NULL REFERENCES Store(store_id) ON DELETE CASCADE,
                                 product_id INT NOT NULL REFERENCES Product(product_id) ON DELETE CASCADE,
                                 stock_quantity INT NOT NULL CHECK (stock_quantity >= 0),
                                 PRIMARY KEY (store_id, product_id)
);

CREATE TABLE IF NOT EXIST Supplier_Catalog (
                                  supplier_id INT NOT NULL REFERENCES Supplier(supplier_id) ON DELETE CASCADE,
                                  product_id INT NOT NULL REFERENCES Product(product_id) ON DELETE CASCADE,
                                  supply_price NUMERIC(10, 2) NOT NULL CHECK (supply_price > 0),
                                  PRIMARY KEY (supplier_id, product_id)
);

CREATE TABLE IF NOT EXIST Sale (
                      sale_id SERIAL PRIMARY KEY,
                      sale_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      seller_id INT NOT NULL REFERENCES Seller(seller_id) ON DELETE RESTRICT,
                      customer_id INT REFERENCES Customer(customer_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXIST Sale_Details (
                              sale_id INT NOT NULL REFERENCES Sale(sale_id) ON DELETE CASCADE,
                              product_id INT NOT NULL REFERENCES Product(product_id) ON DELETE RESTRICT,
                              quantity INT NOT NULL CHECK (quantity > 0),
                              PRIMARY KEY (sale_id, product_id)
);

CREATE TABLE IF NOT EXIST Request (
                         request_id SERIAL PRIMARY KEY,
                         request_date DATE NOT NULL DEFAULT CURRENT_DATE,
                         status VARCHAR(50) NOT NULL DEFAULT 'New',
                         store_id INT NOT NULL REFERENCES Store(store_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXIST Request_Details (
                                 request_id INT NOT NULL REFERENCES Request(request_id) ON DELETE CASCADE,
                                 product_id INT NOT NULL REFERENCES Product(product_id) ON DELETE RESTRICT,
                                 quantity INT NOT NULL CHECK (quantity > 0),
                                 PRIMARY KEY (request_id, product_id)
);

CREATE TABLE IF NOT EXIST Purchase_Order (
                                order_id SERIAL PRIMARY KEY,
                                order_date DATE NOT NULL DEFAULT CURRENT_DATE,
                                status VARCHAR(50) NOT NULL DEFAULT 'Created',
                                supplier_id INT NOT NULL REFERENCES Supplier(supplier_id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXIST Order_Details (
                               order_id INT NOT NULL REFERENCES Purchase_Order(order_id) ON DELETE CASCADE,
                               product_id INT NOT NULL REFERENCES Product(product_id) ON DELETE RESTRICT,
                               quantity INT NOT NULL CHECK (quantity > 0),
                               PRIMARY KEY (order_id, product_id)
);