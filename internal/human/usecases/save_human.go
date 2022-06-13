package usecases

import (
	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/shared/persistence"
)

// SaveHumanUseCase groups structs needed to save human use case.
type SaveHumanUseCase struct {
	repository persistence.HumanRepository
}

// NewSaveHumanUseCase creates a new instance of SaveHumanUseCase.
func NewSaveHumanUseCase(repository persistence.HumanRepository) *SaveHumanUseCase {
	return &SaveHumanUseCase{
		repository: repository,
	}
}

// Save stores a human through a repository.
func (uc *SaveHumanUseCase) Save(human *models.Human) error {
	err := uc.repository.Save(human)
	if err != nil {
		return err
	}
	return nil
}
