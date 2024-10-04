package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"

	db "movies_crud/internal"
)

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range db.MoviesDB {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}