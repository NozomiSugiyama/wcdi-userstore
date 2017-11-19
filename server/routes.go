package server

import (
	"net/http"
	"github.com/wcdi/wcdi-userstore/server/handler"
)

type Route struct {
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"GET",
		"/user",
		handler.UserList,
	},
	Route{
		"POST",
		"/user",
		handler.UserCreate,
	},
	Route{
		"GET",
		"/user/{id}",
		handler.UserShow,
	},
	Route{
		"PUT",
		"/user/{id}",
		handler.UserUpdate,
	},
	Route{
		"DELETE",
		"/user/{id}",
		handler.UserDelete,
	},
}
