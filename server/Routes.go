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

func NewRouter_3000() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes_3000 {
		// register handlers for these URIs. Kind of observer design pattern
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func NewRouter_3001() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes_3001 {
		// register handlers for these URIs. Kind of observer design pattern
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

func NewRouter_3002() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes_3002 {
		// register handlers for these URIs. Kind of observer design pattern
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes_3000 = Routes{
	Route{
		"PutKey_3000",
		"PUT",
		"/keys/{key_id}/{value}",
		PutKey_3000,
	},
	Route{
		"GetKey_3000",
		"GET",
		"/keys/{key_id}",
		GetKey_3000,
	},
	Route{
		"GetAllKeys_3000",
		"GET",
		"/keys",
		GetAllKeys_3000,
	},
}

var routes_3001 = Routes{
	Route{
		"PutKey_3001",
		"PUT",
		"/keys/{key_id}/{value}",
		PutKey_3001,
	},
	Route{
		"GetKey_3001",
		"GET",
		"/keys/{key_id}",
		GetKey_3001,
	},
	Route{
		"GetAllKeys_3001",
		"GET",
		"/keys",
		GetAllKeys_3001,
	},
}

var routes_3002 = Routes{
	Route{
		"PutKey_3002",
		"PUT",
		"/keys/{key_id}/{value}",
		PutKey_3002,
	},
	Route{
		"GetKey_3002",
		"GET",
		"/keys/{key_id}",
		GetKey_3002,
	},
	Route{
		"GetAllKeys_3002",
		"GET",
		"/keys",
		GetAllKeys_3002,
	},
}

var map_3000 = make(map[int]string)
var map_3001 = make(map[int]string)
var map_3002 = make(map[int]string)
