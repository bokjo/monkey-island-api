package handlers

import (
	"net/http"
)

// GetMonkey handles retrieving single Monkey
func GetMonkey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetMonkey HANDLER!"))
}

// CreateMonkey handles single Monkey creation
func CreateMonkey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateMonkey HANDLER!"))
}

// UpdateMonkey handles updating single Monkey
func UpdateMonkey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateMonkey HANDLER!"))
}

// DeleteMonkey handles deleting single Monkey
func DeleteMonkey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteMonkey HANDLER!"))
}

// GetAllMonkeys handles retrieving all the Monkeys
func GetAllMonkeys(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllMonkeys HANDLER!"))
}
