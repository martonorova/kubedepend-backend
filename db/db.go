package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	Client *sql.DB
}

func Get(connString string) (*DB, error) {
	db, err := get(connString)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func get(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
