package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

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

func ColDelete(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)

	err := s.DB(vars["db"]).C(vars["collection"]).DropCollection()
	if err != nil {
		writeError(w, 500, "Error dropping collection")
		return
	}

	w.Header().Add("Content-Length", "0")
	w.WriteHeader(204)
	return
}
