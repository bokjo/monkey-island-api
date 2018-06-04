package handlers

import (
	"database/sql"
	"net/http"
)

// GetDog handles retrieving single dog
func GetDog(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetDog HANDLER!"))
	})
}

// CreateDog handles single dog creation
func CreateDog(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("CreateDog HANDLER!"))
	})
}

// UpdateDog handles updating single dog
func UpdateDog(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateDog HANDLER!"))
	})
}

// DeleteDog handles deleting single dog
func DeleteDog(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeleteDog HANDLER!"))
	})
}

// GetAllDogs handles retrieving all the dogs
func GetAllDogs(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetAllDogs HANDLER!"))
	})
}
