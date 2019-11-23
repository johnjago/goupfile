package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const STATIC_DIR = "/static/"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
