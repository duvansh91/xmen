package usecases

import (
	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/shared/persistence"
)

// GetStatsUseCase groups structs needed to get stats use case.
type GetStatsUseCase struct {
	repository persistence.HumanRepository
}

// NewGetStatsUseCase creates a new instance of GetStatsUseCase
func NewGetStatsUseCase(repository persistence.HumanRepository) *GetStatsUseCase {
	return &GetStatsUseCase{
		repository: repository,
	}
}

// Get gets stats from a repository.
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

	var ratio float32 = 1
	if mutants != 0 {
		ratio = float32(mutants) / float32(humans+mutants)
	}

	stats := &models.Stats{
		CountMutantDNA: int16(mutants),
		CountHumanDNA:  int16(humans),
		Ratio:          float32(ratio),
	}

	return stats, nil
}
