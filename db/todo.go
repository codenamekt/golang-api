package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Name      string        `bson:"name" json:"name"`
	Completed bool          `bson:"completed" json:"completed"`
	Inserted  time.Time     `bson:"inserted" json:"-"`
	Due       time.Time     `bson:"due" json:"due"`
}

type Todos []Todo

/*
type CRRaUD struct {
	Create  func
	Read    func
	ReadAll func
	Update  func
	Delete  func
}
*/

// C
func MongoCreateTodo(t Todo) {
	collection := getCollection()
	err := collection.Insert(t)
	if err != nil {
		panic(err)
	}
}

// R
func MongoReadTodo(t string) Todo {
	collection := getCollection()
	query := bson.M{"_id": bson.ObjectIdHex(t)}
	result := Todo{}
	err := collection.Find(query).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}

// Ra
func MongoReadAllTodo() []Todo {
	collection := getCollection()
	query := bson.M{}
	result := []Todo{}
	err := collection.Find(query).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

// U
func MongoUpdateTodo(id int) string {
	/*
	 * TODO
	 */
	return "something"
}

// D
func MongoDestroyTodo(id int) string {
	/*
	 * TODO
	 */
	return "something"
}
