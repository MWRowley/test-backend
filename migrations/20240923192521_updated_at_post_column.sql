-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts ADD COLUMN updated_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts DROP COLUMN updated_at;
-- +goose StatementEnd
