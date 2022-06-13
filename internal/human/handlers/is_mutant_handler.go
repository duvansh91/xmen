package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/human/usecases"
)

const (
	IsMutantMsg    = "Es un mutante"
	IsNotMutantMsg = "No es un mutante"
)

// GetStatsHandler groups use cases needed to is mutant handler.
type IsMutantHandler struct {
	isMutantUseCase  usecases.ValidateIsMutant
	saveHumanUseCase usecases.SaveHuman
}

// IsMutantRequest defines a is mutant handler request.
type IsMutantRequest struct {
	DNA []string
}

// IsMutantResponse defines a is mutant handler response.
type IsMutantResponse struct {
	Message  string `json:"message"`
	HttpCode int    `json:"code"`
}

// NewIsMutantHandler creates a new instance of IsMutantHandler.
func NewIsMutantHandler(
	isMutantUseCase usecases.ValidateIsMutant,
	saveHumanUseCase usecases.SaveHuman,
) *IsMutantHandler {
	return &IsMutantHandler{
		isMutantUseCase:  isMutantUseCase,
		saveHumanUseCase: saveHumanUseCase,
	}
}

// Handle handles a callback from a request.
func (h *IsMutantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := IsMutantRequest{}
	response := IsMutantResponse{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = "json body malformed"
		response.HttpCode = http.StatusBadRequest
	}

	if len(request.DNA) == 0 {
		response.Message = "invalid DNA"
		response.HttpCode = http.StatusBadRequest
	} else {
		human := &models.Human{
			DNA: request.DNA,
		}

		responseValidation := h.Validate(human)
		response.Message = responseValidation.Message
		response.HttpCode = responseValidation.HttpCode
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.HttpCode)
	json.NewEncoder(w).Encode(response.Message)
}

// Validate validates a mutant through use case.
func (h *IsMutantHandler) Validate(hu *models.Human) *IsMutantResponse {
	response := IsMutantResponse{}

	isMutant, err := h.isMutantUseCase.Validate(hu)
	if err != nil {
		response = IsMutantResponse{
			Message:  err.Error(),
			HttpCode: http.StatusInternalServerError,
		}
	} else {
		hu.IsMutant = false

		response = IsMutantResponse{
			Message:  IsNotMutantMsg,
			HttpCode: http.StatusForbidden,
		}

		if isMutant {
			hu.IsMutant = true

			response = IsMutantResponse{
				Message:  IsMutantMsg,
				HttpCode: http.StatusOK,
			}
		}

		err = h.saveHumanUseCase.Save(hu)
		if err != nil {
			response = IsMutantResponse{
				Message:  err.Error(),
				HttpCode: http.StatusInternalServerError,
			}
		}
	}

	return &response
}
