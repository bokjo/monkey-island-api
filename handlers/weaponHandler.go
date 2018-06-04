package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bokjo/monkey-island-api/model"
	"github.com/gorilla/mux"
)

// WeaponHandler handler struct
type WeaponHandler struct {
	WeaponService model.WeaponService
}

// GetWeapon handles retrieving single Weapon
func (wh *WeaponHandler) GetWeapon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid weapon ID")
		return
	}

	weapon, err := wh.WeaponService.GetWeapon(id)

	if err != nil {

		switch {
		case err == sql.ErrNoRows:
			errorRespond(w, http.StatusNotFound, fmt.Sprintf("Weapon with ID: %d not found", id))
		default:
			errorRespond(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	jsonRespond(w, http.StatusOK, weapon)
}

// CreateWeapon handles single Weapon creation
func (wh *WeaponHandler) CreateWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateWeapon HANDLER!"))
}

// UpdateWeapon handles updating single Weapon
func (wh *WeaponHandler) UpdateWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateWeapon HANDLER!"))
}

// DeleteWeapon handles deleting single Weapon
func (wh *WeaponHandler) DeleteWeapon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteWeapon HANDLER!"))
}

// GetAllWeapons handles retrieving all the Weapons
func (wh *WeaponHandler) GetAllWeapons(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllWeapons HANDLER!"))
}
