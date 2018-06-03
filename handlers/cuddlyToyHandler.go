package handlers

import (
	"net/http"
)

// GetCuddlyToys handles retrieving all the cuddly toys
func GetCuddlyToys(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetCuddlyToys HANDLER!"))
}
