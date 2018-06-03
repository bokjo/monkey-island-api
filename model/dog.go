package model

import (
	"database/sql"
	"errors"
)

// Dog struct - representation of a dog in a DB
type Dog struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EnergyLevel int    `json:"energy_level"`
}

func (d *Dog) getDog(db *sql.DB) error {
	return errors.New("N/A")
}

func (d *Dog) createDog(db *sql.DB) error {
	return errors.New("N/A")
}

func (d *Dog) updateDog(db *sql.DB) error {
	return errors.New("N/A")
}

func (d *Dog) deleteDog(db *sql.DB) error {
	return errors.New("N/A")
}

func (d *Dog) getAllDogs(db *sql.DB) error {
	return errors.New("N/A")
}
