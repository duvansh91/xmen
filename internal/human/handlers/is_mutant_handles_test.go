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
	}

	human := &models.Human{
		DNA: []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
	}

	tests := []struct {
		name     string
		args     args
		funcMock func(mocks mocksProvider, args args)
		mocks    mocksProvider
		want     *Response
	}{
		{
			name: "validate when is mutant",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(true, nil)
			},
			want: &Response{
				Message:  IsMutant,
				HttpCode: http.StatusOK,
			},
		},
		{
			name: "validate when is not mutant",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(false, nil)
			},
			want: &Response{
				Message:  IsNotMutant,
				HttpCode: http.StatusForbidden,
			},
		},
		{
			name: "thrown an error from use case",
			mocks: mocksProvider{
				ValidateIsMutantMock: &mocks.ValidateIsMutant{},
			},
			args: args{
				human: human,
			},
			funcMock: func(mocks mocksProvider, args args) {
				mocks.ValidateIsMutantMock.On("Validate", human).Once().Return(false, errors.New("invalid DNA"))
			},
			want: &Response{
				Message:  "invalid DNA",
				HttpCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.funcMock(tt.mocks, tt.args)
			handler := NewIsMutantHandler(tt.mocks.ValidateIsMutantMock)

			got := handler.HandleValidation(tt.args.human)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.HandleValidation() \nactual = %v \nexpect = %v", got, tt.want)
			}

			tt.mocks.ValidateIsMutantMock.AssertExpectations(t)
		})
	}
}
