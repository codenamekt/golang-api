// Defines handlers for use with mongo db.
// Supported verbs GET, POST, PUT, DELETE.
package handlers

import (
	"gopkg.in/mgo.v2"
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
