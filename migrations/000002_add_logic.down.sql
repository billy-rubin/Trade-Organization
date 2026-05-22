DROP VIEW IF EXISTS v_TurnoverByStoreType;
DROP VIEW IF EXISTS v_ActiveCustomers;
DROP VIEW IF EXISTS v_StoreInventory;

DROP PROCEDURE IF EXISTS sp_ReceiveOrder;
DROP PROCEDURE IF EXISTS sp_TransferProduct;
DROP PROCEDURE IF EXISTS sp_GenerateOrderFromRequest;

DROP TRIGGER IF EXISTS trg_ValidateCustomerStoreType ON Sale;
DROP FUNCTION IF EXISTS tf_ValidateCustomerStoreType;

REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM role_seller, role_purchase_manager, role_director;
REVOKE ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public FROM role_seller, role_purchase_manager, role_director;