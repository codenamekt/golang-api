package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

// globals
var (
	collectionName = "todo"
	currentId      int
	databaseName   = "todo"
	mgoSession     *mgo.Session
	murl           = "localhost:27017"
	todos          Todos
)

func init() {
	//base session
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(murl)
		if err != nil {
			panic(err) // no, not really
		}
	}
	// seed data
	sess := getSession()
	database := session.DB(databaseName)
	collection := database.C(collectionName)
	doc1 := Todo{
		Id:   bson.NewObjectId(),
		Name: "Write presentation"}
	doc2 := Todo{
		Id:   bson.NewObjectId(),
		Name: "Host meetup"}
	RepoCreateTodo(collection, doc1)
	RepoCreateTodo(collection, doc2)
	sess.Close()
}

func getSession() *mgo.Session {
	return mgoSession.Clone()
}

func RepoFindTodo(t string) Todo {
	session := getSession()
	database := session.DB(databaseName)
	collection := database.C(collectionName)
	query := bson.M{"_id": bson.ObjectIdHex(t)}
	result := Todo{}
	err := collection.Find(query).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Todo", result)
	return result
}

// possible race conditions
func RepoCreateTodo(c *mgo.Collection, t Todo) {
	err := c.Insert(t)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}
}

func RepoDestroyTodo(id int) error {
	/*
	   Do stuff here
	*/
	return fmt.Errorf("Could not find Todo")
}
