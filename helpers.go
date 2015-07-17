package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// write to http.ResponseWriter
func writeHttp(w http.ResponseWriter, status int, message string) {
	// write headers
	header := w.Header()
	header.Add("Content-Length", strconv.Itoa(len(message)))
	header.Add("Content-Type", "application/json; charset=UTF-8")

	// write status code
	w.WriteHeader(status)

	// write data
	w.Write([]byte(message))
}

// write response as {"error": message}
func writeError(w http.ResponseWriter, status int, message string) {
	writeHttp(w, status, fmt.Sprintf("{\"error\":\"%s\"}", message))
}
