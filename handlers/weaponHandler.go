package handlers

import (
	"net/http"
)

// GetWeapon handles retrieving single Weapon
func GetWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetWeapon HANDLER!"))
}

// CreateWeapon handles single Weapon creation
func CreateWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateWeapon HANDLER!"))
}

// UpdateWeapon handles updating single Weapon
func UpdateWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateWeapon HANDLER!"))
}

// DeleteWeapon handles deleting single Weapon
func DeleteWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateWeapon HANDLER!"))
}

// GetAllWeapons handles retrieving all the Weapons
func GetAllWeapons(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllWeapons HANDLER!"))
}
