package usecases

import (
	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/shared/persistence"
)

type GetStatsUseCase struct {
	repository persistence.HumanRepository
}

func NewGetStatsUseCase(repository persistence.HumanRepository) *GetStatsUseCase {
	return &GetStatsUseCase{
		repository: repository,
	}
}

func (uc *GetStatsUseCase) Get() (*models.Stats, error) {
	results, err := uc.repository.FindAll()
	if err != nil {
		return nil, err
	}

	mutants := 0
	humans := 0

	for _, human := range results {
		if human.IsMutant {
			mutants += 1
			continue
		}
		humans += 1
	}

	ratio := 1
	if mutants != 0 {
		ratio = mutants / humans
	}

	stats := &models.Stats{
		CountMutantDNA: int16(mutants),
		CountHumanDNA:  int16(humans),
		Ratio:          float32(ratio),
	}

	return stats, nil
}
