package model

import (
	"math/rand"
)

// GhostService - struct
type GhostService struct{}

// GetAllGhosts - generates random number of ghosts <=10
func (gs *GhostService) GetAllGhosts() ([]Ghost, error) {

	numOfGhosts := randomInt(1, 11)

	ghosts := []Ghost{}

	for i := 1; i <= numOfGhosts; i++ {
		ghost := Ghost{
			ID:   i,
			Name: randomString(randomInt(1, 9)),
		}

		ghosts = append(ghosts, ghost)
	}

	return ghosts, nil
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with specified lenght
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
