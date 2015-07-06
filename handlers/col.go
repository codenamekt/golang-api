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
