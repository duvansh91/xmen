package usecases

import "github.com/duvansh91/xmen/internal/human/models"

type ValidateIsMutant interface {
	Validate(human *models.Human) (bool, error)
}

type SaveHuman interface {
	Save(human *models.Human) error
}

type GetStats interface {
	Get() (*models.Stats, error)
}
