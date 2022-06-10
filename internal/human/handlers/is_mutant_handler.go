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
	Message string `json:"message"`
}

func (h *IsMutantHandler) Handle(w http.ResponseWriter, r *http.Request) {

	response := Response{
		Message: IsNotMutant,
	}

	httpCode := http.StatusForbidden

	request := Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response = Response{
			Message: "json body malformed",
		}
	}

	human := models.Human{
		DNA: request.DNA,
	}

	isMutant, err := h.useCase.Validate(&human)

	if isMutant {
		response = Response{
			Message: IsMutant,
		}

		httpCode = http.StatusOK
	}

	if err != nil {
		response = Response{
			Message: err.Error(),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(response)
}

func NewIsMutantHandler(useCase usecases.ValidateIsMutant) *IsMutantHandler {
	return &IsMutantHandler{
		useCase: useCase,
	}
}
