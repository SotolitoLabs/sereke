package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter creates the router based on a route list
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		/* for debugging, delete this content type */
		route.Headers = []string{"Content-type", "(text/html|application/json|text/plain)"}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		//HeadersRegexp(route.Headers...)

	}
	return router
}
