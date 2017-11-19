package server

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/wcdi/wcdi-userstore/server/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func User(w http.ResponseWriter, r *http.Request) {
	users := model.Users{}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}
