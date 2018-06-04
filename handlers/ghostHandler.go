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
	ghosts, err := gh.GhostService.GetAllGhosts()
	if err != nil {
		errorRespond(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonRespond(w, http.StatusOK, ghosts)
}
