package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Port = "8080"

func main() {
	router := httprouter.New()

	server := http.Server{
		Addr:    "localhost:" + Port,
		Handler: router,
	}

	server.ListenAndServe()
}
