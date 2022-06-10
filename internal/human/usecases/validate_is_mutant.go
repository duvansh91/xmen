package usecases

import (
	"errors"
	"strings"

	"github.com/duvansh91/xmen/internal/human/models"
)

type ValidateIsMutantUseCase struct {
}

func NewValidateIsMutantUseCase() *ValidateIsMutantUseCase {
	return &ValidateIsMutantUseCase{}
}

func (uc *ValidateIsMutantUseCase) Validate(h *models.Human) (bool, error) {
	isMutant := false
	n := len(h.DNA) - 1

	// Vertical and horizontal mapping and DNA validation
	for i := 0; i <= n; i++ {
		verticalLastChar := ""
		horizontalLastChar := ""
		verticalCount := 0
		horizontalCount := 0

		for j := 0; j <= n; j++ {
			if !strings.Contains("ATCG", string(h.DNA[j][i])) {
				return false, errors.New("invalid DNA")
			}

			if verticalLastChar == string(h.DNA[j][i]) {
				verticalCount += 1
			} else {
				verticalCount = 0
			}

			if horizontalLastChar == string(h.DNA[i][j]) {
				horizontalCount += 1
			} else {
				horizontalCount = 0
			}

			verticalLastChar = string(h.DNA[j][i])
			horizontalLastChar = string(h.DNA[i][j])

			if verticalCount >= 3 || horizontalCount >= 3 {
				isMutant = true
			}
		}
	}

	// Oblique mapping
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
				currentChar = string(h.DNA[Abs((n-i)+j)][j])
			}

			if i > n {
				currentChar = string(h.DNA[j][Abs(n-i)+j])
			}

			if currentChar == obliqueLastChar {
				obliqueCount += 1
			} else {
				obliqueCount = 0
			}

			if obliqueCount >= 3 {
				isMutant = true
			}

			obliqueLastChar = currentChar
		}
	}

	return isMutant, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
