package app

import (
	"log"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/handlers"
	"github.com/duvansh91/xmen/internal/human/usecases"
	"github.com/duvansh91/xmen/pkg/server"
)

// Run Set up all configurations and instances to run the server
func Run() {
	isMutantUseCase := usecases.NewValidateIsMutantUseCase()
	isMutantHandler := handlers.NewIsMutantHandler(isMutantUseCase)
	isMutantRoute := server.Route{
		Name:    "/mutant/",
		Method:  http.MethodPost,
		Handler: isMutantHandler.Handle,
	}

	routes := []server.Route{
		isMutantRoute,
	}

	s := server.New(routes)

	err := http.ListenAndServe(":8080", s.Router())
	if err != nil {
		log.Fatal(err)
	}
}
