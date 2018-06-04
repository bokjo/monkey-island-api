package model

import (
	"database/sql"
	"errors"
)

// CuddlyToyService - struct
type CuddlyToyService struct {
	DB        *sql.DB
	CuddlyToy *CuddlyToy
}

//GetAllCuddlyToys - retrieve all cuddly toys from DB
func (cts *CuddlyToyService) GetAllCuddlyToys() ([]CuddlyToy, error) {
	return nil, errors.New("N/A")
}
