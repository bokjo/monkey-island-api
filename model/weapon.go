package model

import (
	"database/sql"
	"errors"
)

// WeaponService - struct
type WeaponService struct {
	DB     *sql.DB
	Weapon *Weapon
}

//GetWeapon - retrieves single weapon from the DB
func (ws *WeaponService) GetWeapon(id int64) (*Weapon, error) {
	return nil, errors.New("N/A")
}

// CreateWeapon - creates single weapon in DB
func (ws *WeaponService) CreateWeapon() error {
	return errors.New("N/A")
}

// UpdateWeapon - updates single weapon in DB
func (ws *WeaponService) UpdateWeapon(newWeapon *Weapon) error {
	return errors.New("N/A")
}

// DeleteWeapon - deletes single weapon in DB
func (ws *WeaponService) DeleteWeapon(id int) error {
	return errors.New("N/A")
}

// GetAllWeapons - retrieves all the weapons from DB
func (ws *WeaponService) GetAllWeapons() ([]Weapon, error) {
	return nil, errors.New("N/A")
}
