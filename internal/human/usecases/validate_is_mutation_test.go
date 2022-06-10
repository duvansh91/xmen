package usecases

import (
	"testing"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateIsMutantUseCase_Validate(t *testing.T) {

	type args struct {
		human *models.Human
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		msgErr  string
		want    bool
	}{
		{
			name: "response error when get an invalid DNA",
			args: args{
				human: &models.Human{
					DNA: []string{"XXXXXX", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
				},
			},
			wantErr: true,
			msgErr:  "invalid DNA",
		},
		{
			name: "reponse is mutant when has an oblique sequence",
			args: args{
				human: &models.Human{
					DNA: []string{"ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"},
				},
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "reponse is mutant when has an vertical sequence",
			args: args{
				human: &models.Human{
					DNA: []string{"TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "ACCCTA", "TCACTG"},
				},
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "reponse is mutant when has horizontal sequence match",
			args: args{
				human: &models.Human{
					DNA: []string{"TTGCGA", "CAGTAC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
				},
			},
			wantErr: false,
			want:    true,
		},
		{
			name: "reponse is mutant when has horizontal sequence match",
			args: args{
				human: &models.Human{
					DNA: []string{"TTGCGA", "CAGTAC", "TTATGT", "AGAAGG", "GCCCTA", "TCACTG"},
				},
			},
			wantErr: false,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewValidateIsMutantUseCase()

			result, err := useCase.Validate(tt.args.human)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateIsMutantUseCase.Validate() \nactual = %v \nexpect = %v", err, tt.wantErr)
			}

			if (err != nil) && (tt.msgErr != err.Error()) {
				t.Fatalf("ValidateIsMutantUseCase.Validate() \nactual = %v \nexpect = %v", err, tt.msgErr)
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
