package main

import (
	"database/sql"
	"time"
)

const (
	maxOpenDBConn = 15
	maxIdleDBConn = 5
	maxDBLifetime = 5 * time.Minute
)

func initPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
