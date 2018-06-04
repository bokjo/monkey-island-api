package handlers

import (
	"database/sql"
	"net/http"
)

// GetMonkey handles retrieving single Monkey
func GetMonkey(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetMonkey HANDLER!"))
	})
}

// CreateMonkey handles single Monkey creation
func CreateMonkey(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("CreateMonkey HANDLER!"))
	})

}

// UpdateMonkey handles updating single Monkey
func UpdateMonkey(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateMonkey HANDLER!"))
	})
}

// DeleteMonkey handles deleting single Monkey
func DeleteMonkey(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeleteMonkey HANDLER!"))
	})
}

// GetAllMonkeys handles retrieving all the Monkeys
func GetAllMonkeys(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetAllMonkeys HANDLER!"))
	})
}
