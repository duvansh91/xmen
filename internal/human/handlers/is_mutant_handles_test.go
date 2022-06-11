package handlers

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/human/usecases/mocks"
)

func Test_HandleValidation_HandleValidation(t *testing.T) {
	type args struct {
		human *models.Human
	}

	type mocksProvider struct {
		ValidateIsMutantMock *mocks.ValidateIsMutant
		SaveHumanMock        *mocks.SaveHuman
	}

	human := &models.Human{
		DNA: []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
	}

	tests := []struct {
		name     string
		args     args
		funcMock func(mocks mocksProvider, args args)
		mocks    mocksProvider
		want     *IsMutantResponse
	}{
		{
			name: "validate when is mutant",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
				SaveHumanMock:        &mocks.SaveHuman{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(true, nil)
				mocks.SaveHumanMock.On("Save", human).Once().Return(nil)
			},
			want: &IsMutantResponse{
				Message:  IsMutantMsg,
				HttpCode: http.StatusOK,
			},
		},
		{
			name: "validate when is not mutant",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
				SaveHumanMock:        &mocks.SaveHuman{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(false, nil)
				mocks.SaveHumanMock.On("Save", human).Once().Return(nil)
			},
			want: &IsMutantResponse{
				Message:  IsNotMutantMsg,
				HttpCode: http.StatusForbidden,
			},
		},
		{
			name: "throws an error from validate human use case",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(false, errors.New("invalid DNA"))
			},
			want: &IsMutantResponse{
				Message:  "invalid DNA",
				HttpCode: http.StatusInternalServerError,
			},
		},
		{
			name: "throws an error from save human use case",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
				SaveHumanMock:        &mocks.SaveHuman{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(true, nil)
				mocks.SaveHumanMock.On("Save", human).Once().Return(errors.New("internal server error"))
			},
			want: &IsMutantResponse{
				Message:  "internal server error",
				HttpCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.mocks, tt.args)
			handler := NewIsMutantHandler(tt.mocks.ValidateIsMutantMock, tt.mocks.SaveHumanMock)

			got := handler.HandleValidation(tt.args.human)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.HandleValidation() \nactual = %v \nexpect = %v", got, tt.want)
			}

			tt.mocks.ValidateIsMutantMock.AssertExpectations(t)
		})
	}
}
