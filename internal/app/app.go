package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/handlers"
	"github.com/duvansh91/xmen/internal/human/usecases"
	"github.com/duvansh91/xmen/pkg/mongodb"
	"github.com/duvansh91/xmen/pkg/server"
)

// Run Set up all configurations and instances to run the server
func Run() {
	// Setup database
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	connectionOpts := &mongodb.ConnectionOpts{
		Ctx:      context.TODO(),
		Uri:      config.DBUri,
		Database: config.DBName,
	}

	_, err = mongodb.NewMongoDB(connectionOpts)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Setup routes
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

	// Run server
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), s.Router())
	if err != nil {
		log.Fatal(err)
	}
}
