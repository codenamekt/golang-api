package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var routes = []struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}{
	{"DBIndex", "GET", "/", DBIndex},
	{"DBDelete", "DELETE", "/{db}", DBDelete},
	{"ColIndex", "GET", "/{db}", ColIndex},
	{"ColDelete", "DELETE", "/{db}/{collection}", ColDelete},
	{"DocIndex", "GET", "/{db}/{collection}", DocIndex},
	{"DocPost", "POST", "/{db}/{collection}", DocPost},
	{"DocPut", "PUT", "/{db}/{collection}/{id}", DocPut},
	{"Doc", "GET", "/{db}/{collection}/{id}", DocGet},
	{"DocDelete", "DELETE", "/{db}/{collection}/{id}", DocDelete},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
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
