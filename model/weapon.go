package model

import (
	"database/sql"
)

// WeaponService - struct
type WeaponService struct {
	DB     *sql.DB
	Weapon *Weapon
}

//GetWeapon - retrieves single weapon from the DB
func (ws *WeaponService) GetWeapon(id int) (*Weapon, error) {
	weapon := Weapon{}
	err := ws.DB.QueryRow("SELECT id, name, power_level FROM monkeys WHERE id=$1", id).Scan(&weapon.ID, &weapon.Name, &weapon.PowerLevel)

	return &weapon, err
}

// CreateWeapon - creates single weapon in DB
func (ws *WeaponService) CreateWeapon() error {
	newWeapon := ws.Weapon
	err := ws.DB.QueryRow("INSERT INTO weapons(name, power_level) VALUES($1, $2) RETURNING id", newWeapon.Name, newWeapon.PowerLevel).Scan(&newWeapon.ID)

	return err
}

// UpdateWeapon - updates single weapon in DB
func (ws *WeaponService) UpdateWeapon(id int) error {
	updateWeapon := ws.Weapon
	_, err := ws.DB.Exec("UPDATE weapons SET name=$1, power_level=$2 WHERE id=$3", updateWeapon.Name, updateWeapon.PowerLevel, id)

	return err
}

// DeleteWeapon - deletes single weapon in DB
func (ws *WeaponService) DeleteWeapon(id int) (sql.Result, error) {
	ret, err := ws.DB.Exec("DELETE FROM weapons WHERE id=$1", id)

	return ret, err
}

// GetAllWeapons - retrieves all the weapons from DB
func (ws *WeaponService) GetAllWeapons() ([]Weapon, error) {

	rows, err := ws.DB.Query("SELECT id, name, power_level FROM weapons")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	weapons := []Weapon{}

	for rows.Next() {
		var weapon Weapon

		if err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.PowerLevel); err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}
