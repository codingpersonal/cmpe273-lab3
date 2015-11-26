package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		// register handlers for these URIs. Kind of observer design pattern
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"PutKey",
		"PUT",
		"/keys/{key_id}/{value}",
		PutKey,
	},
	Route{
		"GetKey",
		"GET",
		"/keys/{key_id}",
		GetKey,
	},
}
