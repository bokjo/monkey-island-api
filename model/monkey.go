package model

import (
	"database/sql"
)

// MonkeyService - struct
type MonkeyService struct {
	DB     *sql.DB
	Monkey *Monkey
}

// GetMonkey - retrieves single monkey from DB
func (ms *MonkeyService) GetMonkey(id int) (*Monkey, error) {
	monkey := Monkey{}
	err := ms.DB.QueryRow("SELECT id, name, energy_level FROM monkeys WHERE id=$1", id).Scan(&monkey.ID, &monkey.Name, &monkey.EnergyLevel)

	return &monkey, err
}

// CreateMonkey - creates single monkey in DB
func (ms *MonkeyService) CreateMonkey() error {
	newMonkey := ms.Monkey
	err := ms.DB.QueryRow("INSERT INTO monkeys(name, energy_level) VALUES($1, $2) RETURNING id", newMonkey.Name, newMonkey.EnergyLevel).Scan(&newMonkey.ID)

	return err
}

// UpdateMonkey - updates single monkey in DB
func (ms *MonkeyService) UpdateMonkey(id int) error {
	updateMonkey := ms.Monkey
	_, err := ms.DB.Exec("UPDATE monkeys SET name=$1, energy_level=$2 WHERE id=$3", updateMonkey.Name, updateMonkey.EnergyLevel, id)

	return err
}

// DeleteMonkey - deletes single monkey from DB
func (ms *MonkeyService) DeleteMonkey(id int) (sql.Result, error) {
	ret, err := ms.DB.Exec("DELETE FROM monkeys WHERE id=$1", id)
	return ret, err
}

// GetAllMonkeys - retrieves all the monkeys from DB
func (ms *MonkeyService) GetAllMonkeys() ([]Monkey, error) {

	rows, err := ms.DB.Query("SELECT id, name, energy_level FROM monkeys")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	monkeys := []Monkey{}

	for rows.Next() {
		var monkey Monkey

		if err := rows.Scan(&monkey.ID, &monkey.Name, &monkey.EnergyLevel); err != nil {
			return nil, err
		}
		monkeys = append(monkeys, monkey)
	}

	return monkeys, nil
}
