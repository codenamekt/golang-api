package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestDocIndex(t *testing.T) {
	var body []byte
	resp := httptest.NewRecorder()
	var req *http.Request
	url := "http://localhost/todo/todo"

	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		debug(httputil.DumpRequestOut(req, true))
	}

	router := mux.NewRouter()
	var handler http.Handler

	handler = DocIndex
	router.Methods("GET").
		Path("/{db}/{collection}").
		Name("DocIndex").
		Handler(handler)

	router.ServeHTTP(resp, req)
	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		log.Printf("\n%s", body)
	} else {
		log.Fatalf("ERROR: %s", err)
	}
}
