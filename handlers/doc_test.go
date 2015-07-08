package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocIndex(t *testing.T) {
	var body []byte
	resp := httptest.NewRecorder()
	var req *http.Request
	url := "localhost/todo/todo"

	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		req.Header.Add("Content-Type", "application/json")
	}

	DocIndex(resp, req)
	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		fmt.Printf("%s", body)
	} else {
		log.Fatalf("ERROR: %s", err)
	}
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
