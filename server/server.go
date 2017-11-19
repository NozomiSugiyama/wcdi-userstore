package server

import (
	"net/http"

	"github.com/urfave/negroni"
	"strconv"
)

func StartServer(host string, port int) error {
	router := NewRouter()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	addr := host + ":" + strconv.Itoa(port)

	if err := http.ListenAndServe(addr, n); err != nil {
		return err
	}

	return nil
}
