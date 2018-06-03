package model

import (
	"database/sql"
	"errors"
)

// Weapon struct - representation of a Weapon in a DB
type Weapon struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	PowerLevel int    `json:"power_level"`
}

func (w *Weapon) getWeapon(db *sql.DB) error {
	return errors.New("N/A")
}

func (w *Weapon) createWeapon(db *sql.DB) error {
	return errors.New("N/A")
}

func (w *Weapon) updateWeapon(db *sql.DB) error {
	return errors.New("N/A")
}

func (w *Weapon) deleteWeapon(db *sql.DB) error {
	return errors.New("N/A")
}

func (w *Weapon) getAllWeapons(db *sql.DB) error {
	return errors.New("N/A")
}
