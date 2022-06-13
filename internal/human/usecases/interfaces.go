package usecases

import "github.com/duvansh91/xmen/internal/human/models"

// ValidateIsMutant groups is mutant use case methods.
type ValidateIsMutant interface {
	Validate(human *models.Human) (bool, error)
}

// SaveHuman groups save human use case methods.
type SaveHuman interface {
	Save(human *models.Human) error
}

// GetStats groups get stats use case methods.
type GetStats interface {
	Get() (*models.Stats, error)
}
