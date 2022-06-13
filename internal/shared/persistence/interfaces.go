package persistence

import "github.com/duvansh91/xmen/internal/human/models"

// HumanRepository groups human repository methods.
type HumanRepository interface {
	Save(human *models.Human) error
	FindAll() ([]*models.Human, error)
}
