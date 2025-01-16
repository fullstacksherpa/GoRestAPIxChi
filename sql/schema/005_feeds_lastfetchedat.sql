-- +goose Up
BEGIN;
ALTER TABLE feeds ADD COLUMN IF NOT EXISTS last_fetched_at TIMESTAMP;
COMMIT;

-- +goose Down
BEGIN;
ALTER TABLE feeds DROP COLUMN IF EXISTS last_fetched_at;
COMMIT;