-- +goose Up
-- +goose StatementBegin
ALTER TABLE TICKETS ADD COLUMN title TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE TICKETS DROP COLUMN title;
-- +goose StatementEnd
