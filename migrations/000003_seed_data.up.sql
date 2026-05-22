INSERT INTO Store (area, counter_count, store_type) VALUES
                                                        (2500.00, 30, 'Department_Store'),
                                                        (150.00, 5, 'Shop'),
                                                        (10.00, 1, 'Kiosk'),
                                                        (5.00, 1, 'Tray');

INSERT INTO Department_Store (store_id) VALUES (1);
INSERT INTO Shop (store_id) VALUES (2);
INSERT INTO Kiosk (store_id) VALUES (3);
INSERT INTO Tray (store_id) VALUES (4);

INSERT INTO Section (name, floor, store_id) VALUES
                                                ('Electronics', 1, 1),
                                                ('Home Appliances', 2, 1);

INSERT INTO Hall (name, store_id) VALUES
                                      ('Main Tech Hall', 1),
                                      ('Daily Goods Hall', 2);

INSERT INTO Position (title, salary) VALUES
                                         ('Senior Sales Manager', 60000.00),
                                         ('Consultant', 45000.00),
                                         ('Kiosk Vendor', 35000.00);

INSERT INTO Seller (full_name, position_id, store_id) VALUES
                                                          ('Ivanov Ivan Ivanovich', 1, 1),
                                                          ('Petrov Petr Petrovich', 2, 2),
                                                          ('Sidorova Anna Ivanovna', 3, 3),
                                                          ('Smirnov Alexey Viktorovich', 3, 4);

INSERT INTO Customer (full_name, characteristics) VALUES
                                                      ('Alexandrov Alexander', 'Regular VIP customer, prefers electronics'),
                                                      ('Ekaterina Romanova', 'New customer, interested in home appliances');

INSERT INTO Product (name, current_price) VALUES
                                              ('Smartphone XYZ Pro', 50000.00),
                                              ('Laptop ABC Ultra', 80000.00),
                                              ('Chocolate Bar', 100.00),
                                              ('Mineral Water', 50.00);

INSERT INTO Supplier (name) VALUES
                                ('ElectroOpt LLC'),
                                ('Global Food Supply');

INSERT INTO Supplier_Catalog (supplier_id, product_id, supply_price) VALUES
                                                                         (1, 1, 40000.00),
                                                                         (1, 2, 65000.00),
                                                                         (2, 3, 70.00),
                                                                         (2, 4, 30.00);

INSERT INTO Store_Inventory (store_id, product_id, stock_quantity) VALUES
                                                                       (1, 1, 50),
                                                                       (1, 2, 30),
                                                                       (2, 3, 200),
                                                                       (2, 4, 150),
                                                                       (3, 3, 50),
                                                                       (3, 4, 100),
                                                                       (4, 3, 20),
                                                                       (4, 4, 50);

INSERT INTO Payment (payment_type, amount, payment_date, store_id) VALUES
                                                                       ('Rent', 100000.00, '2026-04-01', 1),
                                                                       ('Utilities', 25000.00, '2026-04-05', 1),
                                                                       ('Rent', 40000.00, '2026-04-01', 2),
                                                                       ('Rent', 5000.00, '2026-04-01', 3);

INSERT INTO Sale (sale_date, seller_id, customer_id) VALUES
                                                         ('2026-04-10 12:30:00', 1, 1),
                                                         ('2026-04-12 14:15:00', 2, 2),
                                                         ('2026-04-15 09:00:00', 3, NULL),
                                                         ('2026-04-15 18:45:00', 4, NULL);

INSERT INTO Sale_Details (sale_id, product_id, quantity) VALUES
                                                             (1, 1, 1),
                                                             (1, 2, 1),
                                                             (2, 3, 5),
                                                             (3, 4, 2),
                                                             (4, 3, 3);

INSERT INTO Request (request_date, status, store_id) VALUES
                                                         ('2026-04-20', 'New', 1),
                                                         ('2026-04-22', 'New', 2);

INSERT INTO Request_Details (request_id, product_id, quantity) VALUES
                                                                   (1, 1, 20),
                                                                   (2, 4, 100);

INSERT INTO Purchase_Order (order_date, status, supplier_id) VALUES
                                                                 ('2026-04-21', 'Created', 1),
                                                                 ('2026-04-23', 'Created', 2);

INSERT INTO Order_Details (order_id, product_id, quantity) VALUES
                                                               (1, 1, 20),
                                                               (2, 4, 100);

INSERT INTO App_Users (username, password_hash, role, seller_id) VALUES
                                                                     ('seller', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'role_seller', 1),
                                                                     ('manager', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'role_purchase_manager', NULL),
                                                                     ('director', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'role_director', NULL);