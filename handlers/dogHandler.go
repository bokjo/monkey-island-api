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

	// TODO: [DRY] - extract common functions to utils/helpers!!!
	// Parse request URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid dog ID")
		return
	}

	dog, err := dh.DogService.GetDog(id)

	if err != nil {

		// TODO: Move to utils/helpers!!!
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
	newDog := &dh.DogService.Dog

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newDog); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for creating new Dog!")
		return
	}

	defer r.Body.Close()

	if err := dh.DogService.CreateDog(); err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, newDog)

}

// UpdateDog handles updating single dog
func (dh *DogHandler) UpdateDog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid dog ID")
		return
	}

	updateDog := &dh.DogService.Dog

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(updateDog); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for updating a Dog!")
		return
	}

	defer r.Body.Close()

	if err := dh.DogService.UpdateDog(id); err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, updateDog)

}

// DeleteDog handles deleting single dog
func (dh *DogHandler) DeleteDog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid dog ID")
		return
	}

	ret, err := dh.DogService.DeleteDog(id)
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	cnt, err := ret.RowsAffected()
	if cnt == 0 {
		errorRespond(w, http.StatusInternalServerError, "NOTHING TO DELETE: The Dog with the specified ID does not exists!")
		return
	}

	jsonRespond(w, http.StatusOK, map[string]string{"Result": "SUCESS: Dog successfully deleted!"})

}

// GetAllDogs handles retrieving all the dogs
func (dh *DogHandler) GetAllDogs(w http.ResponseWriter, r *http.Request) {

	dogs, err := dh.DogService.GetAllDogs()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonRespond(w, http.StatusOK, dogs)

}

// TODO: Move to utils/helpers!!!
// errorRespond - responds with custom JSON error message
func errorRespond(w http.ResponseWriter, statusCode int, message string) {
	jsonRespond(w, statusCode, map[string]string{"error": message})
}

// jsonRespond - returns valid JSON response
func jsonRespond(w http.ResponseWriter, statusCode int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
