package handlers

import (
	"net/http"

	"github.com/bokjo/monkey-island-api/model"
)

// GhostHandler handler struct
type GhostHandler struct {
	GhostService model.GhostService
}

// GetGhosts handles retrieving random number of ghost
func (gh *GhostHandler) GetGhosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetGhosts HANDLER!"))
}
