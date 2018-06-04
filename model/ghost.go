package model

import (
	"errors"
)

// GhostService - struct
type GhostService struct{}

// GetAllGhosts - generates random number of ghosts
func (gs *GhostService) GetAllGhosts() ([]Ghost, error) {
	return nil, errors.New("N/A")
}
