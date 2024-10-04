package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	db "movies_crud/internal"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range db.MoviesDB {
		if item.Id == params["id"] {
			db.MoviesDB = append(db.MoviesDB[:index], db.MoviesDB[index+1:]...)
			json.NewEncoder(w).Encode(db.MoviesDB)

			return
		}
	}
}
