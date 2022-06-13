package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/models"
	"github.com/duvansh91/xmen/internal/human/usecases"
)

// GetStatsHandler groups use cases needed to get stats handler.
type GetStatsHandler struct {
	getStatsUseCase usecases.GetStats
}

// GetStatsResponse defines a get stats handler response.
type GetStatsResponse struct {
	Stats *models.Stats `json:"stats"`
}

// NewGetStatsHandler creates a new instance of GetStatsHandler.
func NewGetStatsHandler(getStatsUseCase usecases.GetStats) *GetStatsHandler {
	return &GetStatsHandler{
		getStatsUseCase: getStatsUseCase,
	}
}

// Handle handles a callback from a request.
func (h *GetStatsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	result, err := h.GetStats()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		content := GetStatsResponse{
			Stats: result,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(content)
	}
}

// GetStats gets stats through use case.
func (h *GetStatsHandler) GetStats() (*models.Stats, error) {
	result, err := h.getStatsUseCase.Get()
	if err != nil {
		return nil, err
	}

	return result, nil
}
