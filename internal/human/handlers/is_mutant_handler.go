package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/human/usecases"
)

const (
	IsMutant    = "Es un mutante"
	IsNotMutant = "No es un mutante"
)

type IsMutantHandler struct {
	useCase usecases.ValidateIsMutant
}

type Request struct {
	DNA []string
}

type Response struct {
	Message  string `json:"message"`
	HttpCode int    `json:"code"`
}

func NewIsMutantHandler(useCase usecases.ValidateIsMutant) *IsMutantHandler {
	return &IsMutantHandler{
		useCase: useCase,
	}
}

func (h *IsMutantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	response := Response{
		Message:  IsNotMutant,
		HttpCode: http.StatusForbidden,
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = "json body malformed"
		response.HttpCode = http.StatusBadRequest
	}

	human := &models.Human{
		DNA: request.DNA,
	}

	responseValidation := h.HandleValidation(human)
	response.Message = responseValidation.Message
	response.HttpCode = responseValidation.HttpCode

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.HttpCode)
	json.NewEncoder(w).Encode(response.Message)
}

func (h *IsMutantHandler) HandleValidation(hu *models.Human) *Response {
	response := Response{}

	response = Response{
		Message:  IsNotMutant,
		HttpCode: http.StatusForbidden,
	}

	isMutant, err := h.useCase.Validate(hu)

	if isMutant {
		response = Response{
			Message:  IsMutant,
			HttpCode: http.StatusOK,
		}
	}

	if err != nil {
		response = Response{
			Message:  err.Error(),
			HttpCode: http.StatusInternalServerError,
		}
	}

	return &response
}
