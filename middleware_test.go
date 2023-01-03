package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieve Request")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware Execute")
	})

	middleware := LogMiddleware{router}

	req := httptest.NewRequest("GET", "http://localhost:"+Port+"/", nil)
	rec := httptest.NewRecorder()

	middleware.serveHTTP(rec, req)

	resp := rec.Result()
	body, err := io.ReadAll(resp.Body)
	errHandler(err)

	assert.Equal(t, "Middleware Execute", string(body))
}
