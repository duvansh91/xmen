package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/human/usecases"
)

type isMutantHandler struct {
	useCase usecases.ValidateIsMutant
}

type Request struct {
	DNA []string
}

type Response struct {
	Message string `json:"message"`
}

func (h *isMutantHandler) Handle(w http.ResponseWriter, r *http.Request) {

	response := Response{
		Message: "NO es mutante!",
	}

	request := Request{}

	json.NewDecoder(r.Body).Decode(&request)
	// if err != nil {
	// 	return nil, errors.New("json deserialization failure")
	// }

	human := models.Human{
		DNA: request.DNA,
	}

	isMutant, _ := h.useCase.Validate(human)
	if isMutant {
		response = Response{
			Message: "SI es mutante!",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Message)
}

func NewIsMutantHandler(useCase usecases.ValidateIsMutant) *isMutantHandler {
	return &isMutantHandler{
		useCase: useCase,
	}
}
