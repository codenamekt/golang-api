package main

import (
	"golang-api/handlers"
	"net/http"
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
		"DBIndex",
		"GET",
		"/",
		handlers.DBIndex,
	},
	Route{
		"ColIndex",
		"GET",
		"/{db}",
		handlers.ColIndex,
	},
	Route{
		"DocIndex",
		"GET",
		"/{db}/{collection}",
		handlers.DocIndex,
	},
	Route{
		"DocPost",
		"POST",
		"/{db}/{collection}",
		handlers.DocPost,
	},
	Route{
		"Doc",
		"GET",
		"/{db}/{collection}/{id}",
		handlers.Doc,
	},
	Route{
		"DocDelete",
		"DELETE",
		"/{db}/{collection}/{id}",
		handlers.DocDelete,
	},
}
