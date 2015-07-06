package handlers

import (
	"fmt"
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
