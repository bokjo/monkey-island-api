package model

import (
	"database/sql"
	"errors"
)

// Ghost struct - representation of a Ghost in a DB
type Ghost struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *Ghost) getAllGhosts(db *sql.DB) error {
	return errors.New("N/A")
}
