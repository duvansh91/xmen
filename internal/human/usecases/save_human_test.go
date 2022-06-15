package usecases

import (
	"errors"
	"testing"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/shared/persistence/mocks"
)

func Test_SaveHumanUseCase_Save(t *testing.T) {
	type args struct {
		human *models.Human
	}

	type mocksProvider struct {
		HumanRepositoryMock *mocks.HumanRepository
	}

	human := &models.Human{
		DNA: []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
	}

	tests := []struct {
		name     string
		args     args
		funcMock func(mocks mocksProvider, args args)
		mocks    mocksProvider
		wantErr  bool
		msgErr   string
	}{
		{
			name: "human respository returns error",
			mocks: mocksProvider{
				HumanRepositoryMock: &mocks.HumanRepository{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.HumanRepositoryMock.On("Save", human).Once().Return(errors.New("internal server error"))
			},
			wantErr: true,
			msgErr:  "internal server error",
		},
		{
			name: "success save human",
			mocks: mocksProvider{
				HumanRepositoryMock: &mocks.HumanRepository{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.HumanRepositoryMock.On("Save", human).Once().Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.mocks, tt.args)
			useCase := NewSaveHumanUseCase(tt.mocks.HumanRepositoryMock)

			err := useCase.Save(tt.args.human)

			if (err != nil) != tt.wantErr {
				t.Fatalf("SaveHumanUseCase.Save() \nactual = %v \nexpect = %v", err, tt.wantErr)
			}

			if (err != nil) && (tt.msgErr != err.Error()) {
				t.Fatalf("SaveHumanUseCase.Save() \nactual = %v \nexpect = %v", err, tt.msgErr)
			}

			tt.mocks.HumanRepositoryMock.AssertExpectations(t)
		})
	}
}
