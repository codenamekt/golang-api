package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func ColIndex(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)

	names, _ := s.DB(vars["db"]).CollectionNames()

	writeHttp(w, 200, fmt.Sprintf("[%s]", strings.Join(names, ",")))
	return
}

func ColDelete(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)

	_ = s.DB(vars["db"]).C(vars["collection"]).DropCollection()

	w.Header().Add("Content-Length", "0")
	w.WriteHeader(204)
	return
}
