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

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemID := p.ByName("itemId")

		text := "Product " + id + " Item " + itemID
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:"+Port+"/products/1/items/1", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	errHandler(err)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterPatternCatchAllParam(t *testing.T) {

	router := httprouter.New()
	router.GET("/iamges/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")

		text := "Image : " + image
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:"+Port+"/images/small/profile.png", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	errHandler(err)

	assert.Equal(t, "404 page not found\n", string(body))
}
