package db

import (
	"gopkg.in/mgo.v2"
)

// globals
var (
	collectionName = "todo"
	currentId      int
	databaseName   = "todo"
	mgoSession     *mgo.Session
	murl           = "localhost:27017"
	todos          Todos
	todo           Todo
)

func main() {
	//base session
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(murl)
		if err != nil {
			panic(err) // no, not really
		}
	}
}

func getSession() *mgo.Session {
	return mgoSession.Clone()
}

func getCollection() *mgo.Collection {
	session := getSession()
	database := session.DB(databaseName)
	collection := database.C(collectionName)
	return collection
}
