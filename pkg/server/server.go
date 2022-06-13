package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

// Server groups server methods.
type Server interface {
	Router() http.Handler
}

// Route groups vars needed to generate a route.
type Route struct {
	Name    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Router gets the http router.
func (a *api) Router() http.Handler {
	return a.router
}

// NewSever creates a new instance of Server.
func NewSever(routes []Route) Server {
	a := &api{}

	r := mux.NewRouter()

	for _, route := range routes {
		r.HandleFunc(route.Name, route.Handler).Methods(route.Method)
	}

	a.router = r
	return a
}
