package model

import (
	"database/sql"
)

// DogService - handles the dog type in DB
type DogService struct {
	DB  *sql.DB
	Dog *Dog
}

// GetDog - retieves single dog from DB
func (ds *DogService) GetDog(id int) (*Dog, error) {

	dog := Dog{}
	err := ds.DB.QueryRow("SELECT id, name, energy_level FROM dogs WHERE id=$1", id).Scan(&dog.ID, &dog.Name, &dog.EnergyLevel)

	return &dog, err
}

// CreateDog - creates single dog in DB
func (ds *DogService) CreateDog() error {
	newDog := ds.Dog
	err := ds.DB.QueryRow("INSERT INTO dogs(name, energy_level) VALUES($1, $2) RETURNING id", newDog.Name, newDog.EnergyLevel).Scan(&newDog.ID)

	return err
}

// UpdateDog - updates single dog in DB
func (ds *DogService) UpdateDog(id int) error {
	updateDog := ds.Dog
	_, err := ds.DB.Exec("UPDATE dogs SET name=$1, energy_level=$2 WHERE id=$3", updateDog.Name, updateDog.EnergyLevel, id)

	return err
}

// DeleteDog - deletes single dog in DB
func (ds *DogService) DeleteDog(id int) (sql.Result, error) {

	ret, err := ds.DB.Exec("DELETE FROM dogs WHERE id=$1", id)
	return ret, err
}

// GetAllDogs -  retrieves all dogs from DB
func (ds *DogService) GetAllDogs() ([]Dog, error) {
	rows, err := ds.DB.Query("SELECT id, name, energy_level FROM dogs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	dogs := []Dog{}

	for rows.Next() {
		var dog Dog

		if err := rows.Scan(&dog.ID, &dog.Name, &dog.EnergyLevel); err != nil {
			return nil, err
		}
		dogs = append(dogs, dog)
	}

	return dogs, nil
}
