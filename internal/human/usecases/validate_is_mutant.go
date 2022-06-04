package usecases

import (
	"github.com/duvansh91/xmen/internal/human/models"
)

type ValidateIsMutantUseCase struct {
}

func NewValidateIsMutantUseCase() *ValidateIsMutantUseCase {
	return &ValidateIsMutantUseCase{}
}

func (uc *ValidateIsMutantUseCase) Validate(h models.Human) (bool, error) {

	return h.DNA[0] == "c", nil
}
