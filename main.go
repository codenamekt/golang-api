package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var session *mgo.Session

// Dial session. Clone session for creating concurrent connections to the db.
func init() {
	var err error
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
}

func main() {
	// serve on 8080
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
