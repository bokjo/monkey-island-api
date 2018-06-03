package handlers

import (
	"net/http"
)

// GetDog handles retrieving single dog
func GetDog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetDog HANDLER!"))
}

// CreateDog handles single dog creation
func CreateDog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateDog HANDLER!"))
}

// UpdateDog handles updating single dog
func UpdateDog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateDog HANDLER!"))
}

// DeleteDog handles deleting single dog
func DeleteDog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateDog HANDLER!"))
}

// GetAllDogs handles retrieving all the dogs
func GetAllDogs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllDogs HANDLER!"))
}
