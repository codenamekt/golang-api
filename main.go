package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)

var session *mgo.Session

// Dial session. Copy session for creating concurrent connections to the db.
// Reads from environment variable MGO. (default="localhost:27017")
func init() {
	var err error
	dial := os.Getenv("MGO")

	session, err = mgo.Dial(dial)
	if err != nil {
		panic(err)
	}
}

// serve on 8080
func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
