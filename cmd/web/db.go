package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxOpenConns    = 25
	maxIdleConns    = 25
	connMaxLifetime = 5 * time.Minute
)

func initPostgresDB(dbAddr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// Database connection is successful
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime) // No limit on connection lifetime

	return db, nil
}
