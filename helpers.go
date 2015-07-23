package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

// rand string generator
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().Unix())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
