package handlers

import (
	"database/sql"
	"net/http"
)

// GetWeapon handles retrieving single Weapon
func GetWeapon(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetWeapon HANDLER!"))
	})
}

// CreateWeapon handles single Weapon creation
func CreateWeapon(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("CreateWeapon HANDLER!"))
	})
}

// UpdateWeapon handles updating single Weapon
func UpdateWeapon(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateWeapon HANDLER!"))
	})
}

// DeleteWeapon handles deleting single Weapon
func DeleteWeapon(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeleteWeapon HANDLER!"))
	})
}

// GetAllWeapons handles retrieving all the Weapons
func GetAllWeapons(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetAllWeapons HANDLER!"))
	})
}
