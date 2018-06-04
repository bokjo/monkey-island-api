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

	w.Write([]byte("GetCuddlyToys HANDLER!"))

}
