package server

import (
	"net/http"

	"github.com/urfave/negroni"
)

func StartServer() error {
	router := NewRouter()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	if err := http.ListenAndServe(":8080", n); err != nil {
		return err
	}

	return nil
}
