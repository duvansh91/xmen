package usecases

import (
	"errors"
	"strings"

	"github.com/duvansh91/xmen/internal/human/models"
)

// SaveHumanUseCase groups structs needed to validate is mutant use case.
type ValidateIsMutantUseCase struct {
}

// NewValidateIsMutantUseCase creates a new instance of ValidateIsMutantUseCase.
func NewValidateIsMutantUseCase() *ValidateIsMutantUseCase {
	return &ValidateIsMutantUseCase{}
}

// Validate checks if a human is mutant based in a DNA.
func (uc *ValidateIsMutantUseCase) Validate(h *models.Human) (bool, error) {
	n := len(h.DNA) - 1

	verticalAndHorizontal, err := verticalAndHorizontalMappingAndValidation(n, h.DNA)
	if err != nil {
		return false, err
	}

	oblique := obliqueMapping(n, h.DNA)

	return verticalAndHorizontal || oblique, nil
}

func verticalAndHorizontalMappingAndValidation(n int, dna []string) (bool, error) {
	for i := 0; i <= n; i++ {
		verticalLastChar := ""
		horizontalLastChar := ""
		verticalCount := 0
		horizontalCount := 0

		for j := 0; j <= n; j++ {
			if !strings.Contains("ATCG", string(dna[j][i])) {
				return false, errors.New("ADN inválido")
			}

			if verticalLastChar == string(dna[j][i]) {
				verticalCount += 1
			} else {
				verticalCount = 0
			}

			if horizontalLastChar == string(dna[i][j]) {
				horizontalCount += 1
			} else {
				horizontalCount = 0
			}

			verticalLastChar = string(dna[j][i])
			horizontalLastChar = string(dna[i][j])

			if verticalCount >= 3 || horizontalCount >= 3 {
				return true, nil
			}
		}
	}

	return false, nil
}

func obliqueMapping(n int, dna []string) bool {
	for i := 0; i <= n*2; i++ {
		obliqueLastChar := ""
		obliqueCount := 0

		limit := i
		if i > n {
			limit = (n * 2) - i
		}

		for j := 0; j <= limit; j++ {
			currentChar := ""

			if i <= n {
				currentChar = string(dna[abs((n-i)+j)][j])
			}

			if i > n {
				currentChar = string(dna[j][abs(n-i)+j])
			}

			if currentChar == obliqueLastChar {
				obliqueCount += 1
			} else {
				obliqueCount = 0
			}

			if obliqueCount >= 3 {
				return true
			}

			obliqueLastChar = currentChar
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
