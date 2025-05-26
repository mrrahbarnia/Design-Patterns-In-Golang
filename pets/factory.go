package pets

import "github.com/mrrahbarnia/Design-Patterns-In-Golang/models"

func NewPet(species string) *models.Pet {
	return &models.Pet{
		Species: species,
	}
}
