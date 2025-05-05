package db

import (
	"database/sql"

	"github.com/ed-henrique/lowdesk/internal/models"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const parameters = "?_pragma=foreign_keys(1)&_pragma=journal_mode(wal)&_pragma=synchronous(normal)"

func New(dbURI string, seed bool) (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite", dbURI+parameters)
	if err != nil {
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		return nil, err
	}

	if seed {
		goose.SetBaseFS(models.Migrations)
		if err := goose.SetDialect("sqlite"); err != nil {
			return nil, err
		}

		if err := goose.Up(dbConn, "migrations"); err != nil {
			return nil, err
		}
	}

	return dbConn, nil
}

func Close(dbConn *sql.DB) error {
	dbTX, err := dbConn.Begin()
	if err != nil {
		return err
	}

	_, err = dbTX.Exec("PRAGMA analysis_limit=400;")
	if err != nil {
		return err
	}

	_, err = dbTX.Exec("PRAGMA optimize;")
	if err != nil {
		return err
	}

	err = dbTX.Commit()
	if err != nil {
		return err
	}

	err = dbConn.Close()
	if err != nil {
		return err
	}

	return nil
}
