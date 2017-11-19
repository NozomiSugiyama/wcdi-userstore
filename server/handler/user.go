package handler

import (
	"encoding/json"
	"net/http"
	"github.com/wcdi/wcdi-userstore/server/model"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	users := model.Users{}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}

	return
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}
