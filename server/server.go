package server

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func StartServer() {
	router := NewRouter()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
