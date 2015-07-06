package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

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

func DBDelete(w http.ResponseWriter, r *http.Request) {
	s := session.New()
	defer s.Close()
	vars := mux.Vars(r)

	err := s.DB(vars["db"]).DropDatabase()
	if err != nil {
		writeError(w, 500, "Error dropping database")
		return
	}

	w.Header().Add("Content-Length", "0")
	w.WriteHeader(204)
	return
}
