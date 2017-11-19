package server

import (
	"net/http"
	"github.com/wcdi/wcdi-userstore/server/handler"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"User",
		"GET",
		"/user",
		handler.UserList,
	},
}