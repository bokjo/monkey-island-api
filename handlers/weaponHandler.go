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

// WeaponHandler handler struct
type WeaponHandler struct {
	WeaponService model.WeaponService
}

// GetWeapon handles retrieving single Weapon
func (wh *WeaponHandler) GetWeapon(w http.ResponseWriter, r *http.Request) {

	// Parse request URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid weapon ID")
		return
	}

	// Fetch weapon from DB
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
	// hold reference to the newWeapon
	newWeapon := &wh.WeaponService.Weapon

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(newWeapon); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for creating new Weapon!")
		return
	}

	defer r.Body.Close()

	err := wh.WeaponService.CreateWeapon()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, newWeapon)
}

// UpdateWeapon handles updating single Weapon
func (wh *WeaponHandler) UpdateWeapon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid weapon ID")
		return
	}

	updateWeapon := &wh.WeaponService.Weapon

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(updateWeapon); err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid request payload for updating a Weapon!")
		return
	}

	defer r.Body.Close()

	if err := wh.WeaponService.UpdateWeapon(id); err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
	}

	jsonRespond(w, http.StatusOK, updateWeapon)
}

// DeleteWeapon handles deleting single Weapon
func (wh *WeaponHandler) DeleteWeapon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorRespond(w, http.StatusBadRequest, "Invalid monkey ID")
		return
	}

	ret, err := wh.WeaponService.DeleteWeapon(id)
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	cnt, err := ret.RowsAffected()
	if cnt == 0 {
		errorRespond(w, http.StatusInternalServerError, "NOTHING TO DELETE: The Weapon with the specified ID does not exists!")
		return
	}

	jsonRespond(w, http.StatusOK, map[string]string{"Result": "SUCESS: Weapon successfully deleted!"})
}

// GetAllWeapons handles retrieving all the Weapons
func (wh *WeaponHandler) GetAllWeapons(w http.ResponseWriter, r *http.Request) {
	weapons, err := wh.WeaponService.GetAllWeapons()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonRespond(w, http.StatusOK, weapons)
}
