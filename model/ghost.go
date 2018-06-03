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

func (g *Ghost) getAllGhosts(db *sql.DB) ([]Ghost, error) {
	return nil, errors.New("N/A")
}
