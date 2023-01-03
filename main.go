package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Port = "8080"

func EerHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println(w, "Welcome HttpRouter")
	})

	server := http.Server{
		Addr:    "localhost:" + Port,
		Handler: router,
	}

	server.ListenAndServe()
}
