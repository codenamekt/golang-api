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
	s := session.Copy()
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

	writeHttp(w, 200, string(res))
	return
}

// Add a new document to a collection if _id omitted. Otherwise, update existing document. id is not required.
func DocPost(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, 500, "Error reading request body")
		return
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, 500, "Error closing request body")
		return
	}

	var req map[string]interface{}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, 500, "Body invalid json")
		return
	}

	if req["_id"] != nil {
		var id bson.ObjectId
		if bson.IsObjectIdHex(req["_id"].(string)) {
			id = bson.ObjectIdHex(req["_id"].(string))
			req["_id"] = id
		} else {
			writeError(w, 400, "Invalid id")
			return
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

		writeHttp(w, 201, string(res))
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

		writeHttp(w, 201, string(res))
	}
	return
}

func DocPut(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		writeError(w, 500, "Error reading request body")
		return
	}
	if err := r.Body.Close(); err != nil {
		writeError(w, 500, "Error closing request body")
		return
	}

	var req map[string]interface{}
	if err := json.Unmarshal(body, &req); err != nil {
		writeError(w, 500, "Body invalid json")
		return
	}

	var id bson.ObjectId
	if bson.IsObjectIdHex(vars["id"]) {
		id = bson.ObjectIdHex(vars["id"])
	} else {
		writeError(w, 400, "Invalid id")
		return
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

	writeHttp(w, 201, string(res))
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
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	var id bson.ObjectId
	if bson.IsObjectIdHex(vars["id"]) {
		id = bson.ObjectIdHex(vars["id"])
	} else {
		writeError(w, 400, "Invalid id")
		return
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

	writeHttp(w, 200, string(res))
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
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	var id bson.ObjectId
	if bson.IsObjectIdHex(vars["id"]) {
		id = bson.ObjectIdHex(vars["id"])
	} else {
		writeError(w, 400, "Invalid id")
		return
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
