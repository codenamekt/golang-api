package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"time"
)

var session *mgo.Session

// Dial session. Copy session for creating concurrent connections to the db.
// Reads from environment variable MGO. (default="localhost:27017")
// maxWait for TTL if connection issue to db
func init() {
	var err error
	dial := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	maxWait := time.Duration(5 * time.Second)

	session, err = mgo.DialWithTimeout(dial, maxWait)
	if err != nil {
		panic(err)
	}
}

// serve on 8080
func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
