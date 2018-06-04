package handlers

import (
	"net/http"

	"github.com/bokjo/monkey-island-api/model"
)

// CuddlyToyHandler handler struct
type CuddlyToyHandler struct {
	CuddlyToyService model.CuddlyToyService
}

// GetCuddlyToys handles retrieving all the cuddly toys
func (cth *CuddlyToyHandler) GetCuddlyToys(w http.ResponseWriter, r *http.Request) {

	cuddlyToys, err := cth.CuddlyToyService.GetAllCuddlyToys()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonRespond(w, http.StatusOK, cuddlyToys)

}
