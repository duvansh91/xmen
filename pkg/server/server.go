package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

type Route struct {
	Name    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func (a *api) Router() http.Handler {

	return a.router
}

func New(routes []Route) Server {
	a := &api{}

	r := mux.NewRouter()

	for _, route := range routes {
		r.HandleFunc(route.Name, route.Handler).Methods(route.Method)
	}

	a.router = r
	return a
}
