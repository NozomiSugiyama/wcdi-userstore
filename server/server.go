package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/user", User)

	log.Fatal(http.ListenAndServe(":8080", router))
}
