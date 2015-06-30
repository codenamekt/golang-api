package main

import (
	"net/http"

	"golang-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoRead",
		"GET",
		"/todos/{todoId}",
		TodoRead,
	},
	Route{
		"TodoReadAll",
		"GET",
		"/todos",
		TodoRead,
	},
}
