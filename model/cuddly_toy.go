package model

import (
	"database/sql"
)

// CuddlyToyService - struct
type CuddlyToyService struct {
	DB        *sql.DB
	CuddlyToy *CuddlyToy
}

//GetAllCuddlyToys - retrieve all cuddly toys from DB
func (cts *CuddlyToyService) GetAllCuddlyToys() ([]CuddlyToy, error) {

	//TODO: CALL GetAllDogs and GetAllMonkeys here
	rows, err := cts.DB.Query("SELECT id, name, energy_level FROM monkeys UNION SELECT id, name, energy_level FROM dogs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cuddlyToys := []CuddlyToy{}

	for rows.Next() {
		var cuddlyToy CuddlyToy

		if err := rows.Scan(&cuddlyToy.ID, &cuddlyToy.Name, &cuddlyToy.EnergyLevel); err != nil {
			return nil, err
		}
		cuddlyToys = append(cuddlyToys, cuddlyToy)
	}

	return cuddlyToys, nil
}
