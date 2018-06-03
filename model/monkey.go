package model

import (
	"database/sql"
	"errors"
)

// Monkey struct - representation of a Monkey in a DB
type Monkey struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EnergyLevel int    `json:"energy_level"`
}

func (m *Monkey) getMonkey(db *sql.DB) error {
	return errors.New("N/A")
}

func (m *Monkey) createMonkey(db *sql.DB) error {
	return errors.New("N/A")
}

func (m *Monkey) updateMonkey(db *sql.DB) error {
	return errors.New("N/A")
}

func (m *Monkey) deleteMonkey(db *sql.DB) error {
	return errors.New("N/A")
}

func (m *Monkey) getAllMonkeys(db *sql.DB) ([]Monkey, error) {
	return nil, errors.New("N/A")
}
