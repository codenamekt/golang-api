package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strconv"
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

	router := NewRouter()

	router.ServeHTTP(resp, req)
	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		log.Println("Return:", strconv.Itoa(resp.Code))
		for key, value := range resp.Header() {
			log.Println(key, value)
		}
		log.Println("Body:", string(body))
	} else {
		log.Fatalf("ERROR: %s", err)
	}
}
