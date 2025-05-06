-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS TICKETS (
    ID INTEGER PRIMARY KEY,
    CONTENT TEXT NOT NULL,
    CREATED_AT TEXT DEFAULT (DATETIME('now')) NOT NULL
) STRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE TICKETS;
-- +goose StatementEnd
