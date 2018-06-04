package model

import (
	"database/sql"
	"errors"
)

// DogService - handles the dog type in DB
type DogService struct {
	DB  *sql.DB
	Dog *Dog
}

// GetDog - retieves single dog from DB
func (ds *DogService) GetDog(id int64) (*Dog, error) {

	dog := Dog{}
	err := ds.DB.QueryRow("SELECT id, name, energy_level FROM dogs WHERE id=$1", id).Scan(&dog.ID, &dog.Name, &dog.EnergyLevel)

	return &dog, err
}

// CreateDog - creates single dog in DB
func (ds *DogService) CreateDog() error {
	return errors.New("N/A")
}

// UpdateDog - updates single dog in DB
func (ds *DogService) UpdateDog(newDog *Dog) error {
	return errors.New("N/A")
}

// DeleteDog - deletes single dog in DB
func (ds *DogService) DeleteDog(id int) error {
	return errors.New("N/A")
}

// GetAllDogs -  retrieves all dogs from DB
func (ds *DogService) GetAllDogs() error {
	return errors.New("N/A")
}
