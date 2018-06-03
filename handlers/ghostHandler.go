package handlers

import (
	"net/http"
)

// GetGhosts handles retrieving random number of ghost
func GetGhosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetGhosts HANDLER!"))
}
