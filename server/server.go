package server

import (
	"log"
	"net/http"
)

func StartServer() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
