package main

import (
    "time"
    "gopkg.in/mgo.v2/bson"
)

type Todo struct {
    Id        bson.ObjectId `bson:"_id,omitempty" json:"-"`
    Name      string        `bson:"name" json:"name"`
    Completed bool          `bson:"completed" json:"completed"`
    Inserted  time.Time     `bson:"inserted" json:"-"`
    Due       time.Time     `bson:"due" json:"due"`
}

type Todos []Todo
