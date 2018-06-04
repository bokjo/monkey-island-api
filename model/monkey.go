package model

import (
	"database/sql"
	"errors"
)

// MonkeyService - struct
type MonkeyService struct {
	DB     *sql.DB
	Monkey *Monkey
}

// GetMonkey - retrieves single monkey from DB
func (ms *MonkeyService) GetMonkey(id int64) (*Monkey, error) {
	monkey := Monkey{}
	err := ms.DB.QueryRow("SELECT id, name, energy_level FROM monkeys WHERE id=$1", id).Scan(&monkey.ID, &monkey.Name, &monkey.EnergyLevel)

	return &monkey, err
}

// CreateMonkey - creates single monkey in DB
func (ms *MonkeyService) CreateMonkey() error {
	return errors.New("N/A")
}

// UpdateMonkey - updates single monkey in DB
func (ms *MonkeyService) UpdateMonkey(newMonkey *Monkey) error {
	return errors.New("N/A")
}

// DeleteMonkey - deletes single monkey from DB
func (ms *MonkeyService) DeleteMonkey(id int) error {
	return errors.New("N/A")
}

// GetAllMonkeys - retrieves all the monkeys from DB
func (ms *MonkeyService) GetAllMonkeys() ([]Monkey, error) {
	return nil, errors.New("N/A")
}
