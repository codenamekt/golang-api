package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// List available documents in a collection.
func DocIndex(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	var out []map[string]interface{}
	err := c.Find(nil).All(&out)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		writeError(w, 500, "Error getting all documents")
		return
	}

	res, err := json.Marshal(out)
	if err != nil {
		writeError(w, 500, "Error stringifying response")
		return
	}

	writeJSON(w, 200, string(res))
	return
}

// Add a new document to a collection if _id omitted. Otherwise, update existing document. id is not required.
func DocPost(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, 500, "Error reading request body")
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, 500, "Error closing request body")
	}

	var req map[string]interface{}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, 500, "Body invalid json")
	}

	if req["_id"] != nil {
		id := bson.ObjectIdHex(req["_id"].(string))
		if !id.Valid() {
			writeError(w, 400, "Invalid id")
		} else {
			delete(req, "_id")
		}

		_, err := c.UpsertId(id, req)
		if err != nil {
			writeError(w, 500, "Error upserting")
			return
		}

		res, err := json.Marshal(req)
		if err != nil {
			writeError(w, 500, "Error stringifying query result")
			return
		}

		writeJSON(w, 201, string(res))
	} else {
		req["_id"] = bson.NewObjectId()

		err := c.Insert(req)
		if err != nil {
			writeError(w, 500, "Document insertion failed")
			return
		}

		res, err := json.Marshal(req)
		if err != nil {
			writeError(w, 500, "Error stringifying query result")
			return
		}

		writeJSON(w, 201, string(res))
	}
	return
}

func DocPut(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, 500, "Error reading request body")
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, 500, "Error closing request body")
	}

	var req map[string]interface{}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, 500, "Body invalid json")
	}

	id := bson.ObjectIdHex(vars["id"])
	if !id.Valid() {
		writeError(w, 400, "Invalid id")
	} else {
		req["_id"] = id
	}

	_, err = c.UpsertId(id, req)
	if err != nil {
		writeError(w, 500, "Error upserting")
		return
	}

	res, err := json.Marshal(req)
	if err != nil {
		writeError(w, 500, "Error stringifying query result")
		return
	}

	writeJSON(w, 201, string(res))
	return
}

// Retrieve document from database.
//
// Responses
//
// 200 Success.
//
// 400 Invalid id.
//
// 500 Error counting results. Error getting document by id. Error stringifying query result.
//
// 404 Document not found.
//
func DocGet(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	id := bson.ObjectIdHex(vars["id"])
	if !id.Valid() {
		writeError(w, 400, "Invalid id")
	}

	q := c.Find(bson.M{"_id": id})
	n, err := q.Count()
	if err != nil {
		writeError(w, 500, "Error counting results")
		return
	} else if n == 0 {
		writeError(w, 404, "Document not found")
		return
	}

	var out map[string]interface{}
	err = q.One(&out)
	if err != nil {
		writeError(w, 500, "Error getting document by id")
		return
	}

	res, err := json.Marshal(out)
	if err != nil {
		writeError(w, 500, "Error stringifying query result")
		return
	}

	writeJSON(w, 200, string(res))
	return
}

// Deletes a document from the database.
//
// Responses
//
// 204 Success.
//
// 400 Invalid id.
//
// 500 Error counting results. Error communicating with database. Error removing item from database.
//
// 404 Document not found.
//
func DocDelete(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	id := bson.ObjectIdHex(vars["id"])
	if !id.Valid() {
		writeError(w, 400, "Invalid id")
	}

	q := c.Find(bson.M{"_id": id})
	n, err := q.Count()
	if err != nil {
		writeError(w, 500, "Error counting results")
		return
	} else if n == 0 {
		writeError(w, 404, "Document not found")
		return
	}

	obj := bson.M{"_id": id}

	n, err = c.Find(obj).Count()
	if err != nil {
		writeError(w, 500, "Error communicating with database")
		return
	}
	if n == 0 {
		writeError(w, 404, "Document not found")
		return
	}

	err = c.Remove(obj)
	if err != nil {
		writeError(w, 500, "Error removing item from database")
		return
	}

	w.Header().Add("Content-Length", "0")
	w.WriteHeader(204)
	return
}
