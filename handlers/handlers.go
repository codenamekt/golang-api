package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var session *mgo.Session

func init() {
	var err error
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
}

func writeJSON(w http.ResponseWriter, status int, message string) {
	// write headers
	header := w.Header()
	header.Add("Content-Length", strconv.Itoa(len(message)))
	header.Add("Content-Type", "application/json; charset=UTF-8")

	// write status code
	w.WriteHeader(status)

	// write data
	w.Write([]byte(message))
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, fmt.Sprintf("{\"error\":\"%s\"}", message))
}

func DBIndex(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()

	names, err := s.DatabaseNames()
	if err != nil {
		writeError(w, 500, "Error getting database names")
		return
	}

	writeJSON(w, 200, fmt.Sprintf("[%s]", strings.Join(names, ",")))
	return
}

func ColIndex(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)

	names, err := s.DB(vars["db"]).CollectionNames()
	if err != nil {
		writeError(w, 500, "Error getting collection names")
		return
	}

	writeJSON(w, 200, fmt.Sprintf("[%s]", strings.Join(names, ",")))
	return
}

func DocIndex(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)
	c := s.DB(vars["db"]).C(vars["collection"])

	var out []map[string]interface{}
	err := c.Find(nil).All(&out)
	if err != nil {
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
		writeError(w, 404, "The id provided was not found in this database. Either the document has been deleted already or it never existed.")
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
