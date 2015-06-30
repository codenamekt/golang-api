package main

import (
	"log"
	"net/http"
)

func main() {
	// serve on 8080
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
