package usecases

import (
	"errors"
	"testing"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/shared/persistence/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetStatsUseCase_Get(t *testing.T) {
	type mocksProvider struct {
		HumanRepositoryMock *mocks.HumanRepository
	}

	tests := []struct {
		name     string
		funcMock func(mocks mocksProvider)
		mocks    mocksProvider
		wantErr  bool
		msgErr   string
		want     *models.Stats
	}{
		{
			name: "human respository returns error",
			mocks: mocksProvider{
				HumanRepositoryMock: &mocks.HumanRepository{},
			},
			funcMock: func(mocks mocksProvider) {
				mocks.HumanRepositoryMock.On("FindAll").Once().Return(nil, errors.New("internal server error"))
			},
			wantErr: true,
			msgErr:  "internal server error",
			want:    nil,
		},
		{
			name: "success get stats",
			mocks: mocksProvider{
				HumanRepositoryMock: &mocks.HumanRepository{},
			},
			funcMock: func(mocks mocksProvider) {
				result := []*models.Human{
					{
						DNA:      []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
						IsMutant: false,
					},
					{
						DNA:      []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
						IsMutant: false,
					},
					{
						DNA:      []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
						IsMutant: true,
					},
				}
				mocks.HumanRepositoryMock.On("FindAll").Once().Return(result, nil)
			},
			wantErr: false,
			want: &models.Stats{
				CountMutantDNA: int16(1),
				CountHumanDNA:  int16(2),
				Ratio:          float64(0.3),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.mocks)
			useCase := NewGetStatsUseCase(tt.mocks.HumanRepositoryMock)

			got, err := useCase.Get()

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetStatsUseCase.Get() \nactual = %v \nexpect = %v", err, tt.wantErr)
			}

			if (err != nil) && (tt.msgErr != err.Error()) {
				t.Fatalf("GetStatsUseCase.Get() \nactual = %v \nexpect = %v", err, tt.msgErr)
			}

			assert.Equal(t, tt.want, got)

			tt.mocks.HumanRepositoryMock.AssertExpectations(t)
		})
	}
}
