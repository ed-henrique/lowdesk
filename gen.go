package gen

//go:generate goose -dir internal/models/migrations sqlite3 db.sqlite3 up
//go:generate sqlc -f internal/models/sqlc.yml generate
