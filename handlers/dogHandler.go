package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bokjo/monkey-island-api/model"
)

// DogHandler handler struct
type DogHandler struct {
	DogService model.DogService
}

// GetDog handles retrieving single dog
func (dh *DogHandler) GetDog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid dog ID")
		return
	}

	dog, err := dh.DogService.GetDog(id)

	if err != nil {

		switch {
		case err == sql.ErrNoRows:
			errorRespond(w, http.StatusNotFound, fmt.Sprintf("Dog with ID: %d not found", id))
		default:
			errorRespond(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	jsonRespond(w, http.StatusOK, dog)

}

// CreateDog handles single dog creation
func (dh *DogHandler) CreateDog(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("CreateDog HANDLER!"))

}

// UpdateDog handles updating single dog
func (dh *DogHandler) UpdateDog(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("UpdateDog HANDLER!"))

}

// DeleteDog handles deleting single dog
func (dh *DogHandler) DeleteDog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteDog HANDLER!"))

}

// GetAllDogs handles retrieving all the dogs
func (dh *DogHandler) GetAllDogs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllDogs HANDLER!"))

}

func errorRespond(w http.ResponseWriter, statusCode int, message string) {
	jsonRespond(w, statusCode, map[string]string{"error": message})
}

func jsonRespond(w http.ResponseWriter, statusCode int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
