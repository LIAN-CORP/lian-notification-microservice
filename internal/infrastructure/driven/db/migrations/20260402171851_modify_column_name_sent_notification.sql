-- +goose Up
ALTER TABLE notifications RENAME COLUMN snet_at TO sent_at;

-- +goose Down
ALTER TABLE notifications RENAME COLUMN sent_at TO snet_at;
