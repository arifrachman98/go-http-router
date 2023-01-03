package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	dir, err := fs.Sub(resources, "resources")
	errHandler(err)
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest("GET", "http://localhost:"+Port+"/files/hello.txt", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	errHandler(err)

	assert.Equal(t, "Hello HttpRouter", string(body))
}
