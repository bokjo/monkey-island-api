package handlers

import (
	"database/sql"
	"net/http"
)

// GetCuddlyToys handles retrieving all the cuddly toys
func GetCuddlyToys(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GetCuddlyToys HANDLER!"))
	})
}
