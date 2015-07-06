package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, status int, message string) {
	// write headers
	header := w.Header()
	header.Add("Content-Length", strconv.Itoa(len(message)))
	header.Add("Content-Type", "application/json; charset=UTF-8")

	// write status code
	w.WriteHeader(status)

	// write data
	w.Write([]byte(message))
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, fmt.Sprintf("{\"error\":\"%s\"}", message))
}
