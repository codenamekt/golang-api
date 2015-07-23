package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func DBIndex(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()

	names, _ := s.DatabaseNames()

	writeHttp(w, 200, fmt.Sprintf("[%s]", strings.Join(names, ",")))
	return
}

func DBDelete(w http.ResponseWriter, r *http.Request) {
	s := session.Copy()
	defer s.Close()
	vars := mux.Vars(r)

	_ = s.DB(vars["db"]).DropDatabase()

	w.Header().Add("Content-Length", "0")
	w.WriteHeader(204)
	return
}
