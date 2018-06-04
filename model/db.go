package model

import (
	"database/sql"

	//Postgres driver
	_ "github.com/lib/pq"
)

// DBConnect creates new DB connection handler
func DBConnect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
