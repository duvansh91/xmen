package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/duvansh91/xmen/internal/human/handlers"
	"github.com/duvansh91/xmen/internal/human/usecases"
	"github.com/duvansh91/xmen/internal/shared/persistence"
	"github.com/duvansh91/xmen/pkg/mongodb"
	"github.com/duvansh91/xmen/pkg/server"
)

// Run set up all configurations and instances to run the server
func Run() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	connectionOpts := &mongodb.ConnectionOpts{
		Ctx:      context.TODO(),
		Uri:      config.DBUri,
		Database: config.DBName,
	}

	client, err := mongodb.NewMongoDB(connectionOpts)
	if err != nil {
		log.Fatal(err.Error())
	}

	humanRepository := persistence.NewHumanMongoDBRepository(client)

	isMutantUseCase := usecases.NewValidateIsMutantUseCase()
	saveHumanUsecase := usecases.NewSaveHumanUseCase(humanRepository)
	isMutantHandler := handlers.NewIsMutantHandler(isMutantUseCase, saveHumanUsecase)
	isMutantRoute := server.Route{
		Name:    "/mutant/",
		Method:  http.MethodPost,
		Handler: isMutantHandler.Handle,
	}

	getStatsUsecase := usecases.NewGetStatsUseCase(humanRepository)
	getStatsHandler := handlers.NewGetStatsHandler(getStatsUsecase)
	getStatsRoute := server.Route{
		Name:    "/stats/",
		Method:  http.MethodGet,
		Handler: getStatsHandler.Handle,
	}

	routes := []server.Route{
		isMutantRoute,
		getStatsRoute,
	}

	s := server.NewSever(routes)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), s.Router())
	if err != nil {
		log.Fatal(err)
	}
}
