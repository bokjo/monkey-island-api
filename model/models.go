package model

//CuddlyToy model
type CuddlyToy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EnergyLevel int    `json:"energy_level"`
}

// Dog model
type Dog struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EnergyLevel int    `json:"energy_level"`
}

// Monkey model
type Monkey struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	EnergyLevel int    `json:"energy_level"`
}

// Weapon model
type Weapon struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	PowerLevel int    `json:"power_level"`
}

// Ghost model
type Ghost struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
