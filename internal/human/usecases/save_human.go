package usecases

import (
	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/shared/persistence"
)

type SaveHumanUseCase struct {
	repository persistence.HumanRepository
}

func NewSaveHumanUseCase(repository persistence.HumanRepository) *SaveHumanUseCase {
	return &SaveHumanUseCase{
		repository: repository,
	}
}

func (uc *SaveHumanUseCase) Save(human *models.Human) error {
	err := uc.repository.Save(human)
	if err != nil {
		return err
	}
	return nil
}
