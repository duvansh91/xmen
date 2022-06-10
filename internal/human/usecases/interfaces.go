package usecases

import "github.com/duvansh91/xmen/internal/human/models"

type ValidateIsMutant interface {
	Validate(h *models.Human) (bool, error)
}
