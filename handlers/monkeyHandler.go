package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bokjo/monkey-island-api/model"
	"github.com/gorilla/mux"
)

// MonkeyHandler handler struct
type MonkeyHandler struct {
	MonkeyService model.MonkeyService
}

// GetMonkey handles retrieving single Monkey
func (mh *MonkeyHandler) GetMonkey(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// id, err := strconv.ParseInt(vars["id"], 10, 64)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid monkey ID")
		return
	}
	//mh.MonkeyService.Monkey = &model.Monkey{ID: id}

	monkey, err := mh.MonkeyService.GetMonkey(id)

	if err != nil {

		switch {
		case err == sql.ErrNoRows:
			errorRespond(w, http.StatusNotFound, fmt.Sprintf("Monkey with ID: %d not found", id))
		default:
			errorRespond(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	jsonRespond(w, http.StatusOK, monkey)

}

// CreateMonkey handles single Monkey creation
func (mh *MonkeyHandler) CreateMonkey(w http.ResponseWriter, r *http.Request) {

	// hold reference to the newMonkey
	newMonkey := &mh.MonkeyService.Monkey

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(newMonkey); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for creating new Monkey!")
		return
	}

	defer r.Body.Close()

	err := mh.MonkeyService.CreateMonkey()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, newMonkey)

}

// UpdateMonkey handles updating single Monkey
func (mh *MonkeyHandler) UpdateMonkey(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid monkey ID")
		return
	}

	updateMonkey := &mh.MonkeyService.Monkey

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(updateMonkey); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for updating a Monkey!")
		return
	}

	defer r.Body.Close()

	if err := mh.MonkeyService.UpdateMonkey(id); err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, updateMonkey)

}

// DeleteMonkey handles deleting single Monkey
func (mh *MonkeyHandler) DeleteMonkey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid monkey ID")
		return
	}

	ret, err := mh.MonkeyService.DeleteMonkey(id)
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	cnt, err := ret.RowsAffected()
	if cnt == 0 {
		errorRespond(w, http.StatusInternalServerError, "NOTHING TO DELETE: The Monkey with the specified ID does not exists!")
		return
	}

	jsonRespond(w, http.StatusOK, map[string]string{"Result": "SUCESS: Monkey successfully deleted!"})

}

// GetAllMonkeys handles retrieving all the Monkeys
func (mh *MonkeyHandler) GetAllMonkeys(w http.ResponseWriter, r *http.Request) {

	monkeys, err := mh.MonkeyService.GetAllMonkeys()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonRespond(w, http.StatusOK, monkeys)

}
