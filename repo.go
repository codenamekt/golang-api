package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "os"
)

var (
    currentId int
    murl = "localhost:27017"
    databaseName = "todo"
    mgoSession     *mgo.Session
    todos Todos
)

func getSession () *mgo.Session {
    if mgoSession == nil {
        var err error
        mgoSession, err = mgo.Dial(murl)
        if err != nil {
             panic(err) // no, not really
        }
    }
    return mgoSession.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
    session := getSession()
    defer session.Close()
    c := session.DB(databaseName).C(collection)
    return s(c)
}

func SearchTodo (q interface{}, skip int, limit int) (searchResults []Todo, searchErr string) {
    searchErr     = ""
    searchResults = []Todo{}
    query := func(c *mgo.Collection) error {
        fn := c.Find(q).Skip(skip).Limit(limit).All(&searchResults)
        if limit < 0 {
            fn = c.Find(q).Skip(skip).All(&searchResults)
        }
        return fn
    }
    err := func() error {
        return withCollection("todo", query)
    }
    if err != nil {
        searchErr = "Database Error"
    }
    return
}

// Give us some seed data
func init() {
    sess := getSession()
    c := sess.DB("todo").C("todo")
    doc1 := Todo{
        Id: bson.NewObjectId(),
        Name: "Write presentation"}
    doc2 := Todo{
        Id: bson.NewObjectId(),
        Name: "Host meetup"}
    RepoCreateTodo(c, doc1)
    RepoCreateTodo(c, doc2)
    sess.Close()
}

func GetTodoById (Id string) (searchResults []Todo, searchErr string) {
    searchResults, searchErr = SearchTodo(bson.M{"_id": bson.ObjectIdHex(Id)}, 0, 0)
    return
}

func RepoFindTodo(Id string) []Todo {
    todo, err := GetTodoById(Id)
    if err != "" {
        fmt.Printf("Can't find document: %d %v\n", Id, err)
        return []Todo{}
    }
    return todo
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(c *mgo.Collection, t Todo) {
    err := c.Insert(t)
    if err != nil {
        fmt.Printf("Can't insert document: %v\n", err)
        os.Exit(1)
    }
}

func RepoDestroyTodo(id int) error {
    /*
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    */
    return fmt.Errorf("Could not find Todo")
}
