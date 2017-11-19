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
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
}
