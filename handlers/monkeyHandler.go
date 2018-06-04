package handlers

import (
	"database/sql"
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
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid monkey ID")
		return
	}

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

	w.Write([]byte("CreateMonkey HANDLER!"))

}

// UpdateMonkey handles updating single Monkey
func (mh *MonkeyHandler) UpdateMonkey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateMonkey HANDLER!"))

}

// DeleteMonkey handles deleting single Monkey
func (mh *MonkeyHandler) DeleteMonkey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteMonkey HANDLER!"))

}

// GetAllMonkeys handles retrieving all the Monkeys
func (mh *MonkeyHandler) GetAllMonkeys(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("GetAllMonkeys HANDLER!"))

}
