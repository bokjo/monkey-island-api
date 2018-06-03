package model

import (
	"database/sql"
	"errors"
)

// CuddlyToy struct
type CuddlyToy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EnergyLevel int    `json:"energy_level"`
}

func (ct *CuddlyToy) getAllCuddlyToys(db *sql.DB) error {
	return errors.New("N/A")
}
