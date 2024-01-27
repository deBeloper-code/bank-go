ALTER TABLE transfers DROP CONSTRAINT IF EXISTS transfers_from_account_id_fkey;
ALTER TABLE transfers DROP CONSTRAINT IF EXISTS transfers_to_account_id_fkey;
ALTER TABLE accounts DROP CONSTRAINT IF EXISTS accounts_owner_id_fkey;

DROP INDEX IF EXISTS idx_username;
DROP INDEX IF EXISTS idx_email;
DROP INDEX IF EXISTS idx_owner_id;
DROP INDEX IF EXISTS idx_from_account_id;
DROP INDEX IF EXISTS idx_to_account_id;
DROP INDEX IF EXISTS idx_transfer_accounts;

COMMENT ON COLUMN transfers.amount IS NULL;

DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS users;





